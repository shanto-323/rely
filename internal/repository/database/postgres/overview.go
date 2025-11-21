package postgres

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/shanto-323/rely/model"
	"github.com/shanto-323/rely/model/dto"
)

func (d *DB) StudentAttendanceOverview(ctx context.Context, studentId int) (*dto.StudentAttendanceOverview, error) {
	query := `
        WITH student_info AS (
            SELECT
                id,
                fullname,
                department,
                shift,
                semester,
                section
            FROM students
            WHERE student_id = @student_id
        ),

        class_counts AS (
            SELECT 
                COUNT(*) AS total_classes,
                SUM(CASE WHEN ar.student_id IS NOT NULL THEN 1 ELSE 0 END) AS present
            FROM attendance_sessions s
            LEFT JOIN attendance_records ar 
                ON ar.session_id = s.id 
                AND ar.student_id = @student_id
            CROSS JOIN student_info si
            WHERE 
                s.department = si.department
                AND s.shift = si.shift
                AND s.semester = si.semester
                AND s.section = si.section
        ),

        limited_sessions AS (
            SELECT 
                s.id,
                s.teacher_id,
                t.fullname AS teacher_name,
                t.email AS teacher_email,
                t.phone AS teacher_phone,
                s.subject_code,
                s.created_at,
                ar.student_id IS NOT NULL AS present
            FROM attendance_sessions s
            LEFT JOIN attendance_records ar 
                ON ar.session_id = s.id
                AND ar.student_id = @student_id
            JOIN teachers t ON t.id = s.teacher_id
            CROSS JOIN student_info si
            WHERE 
                s.department = si.department
                AND s.shift = si.shift
                AND s.semester = si.semester
                AND s.section = si.section
            ORDER BY s.created_at DESC
            LIMIT 10
        ),

        last_sessions AS (
            SELECT json_agg(
                json_build_object(
                    'session_id', id,
                    'teacher', json_build_object(
                        'id', teacher_id,
                        'fullname', teacher_name,
                        'email', teacher_email,
                        'phone', teacher_phone
                    ),
                    'subject_code', subject_code,
                    'created_at', created_at,
                    'present', present
                )
                ORDER BY created_at DESC
            ) AS sessions
            FROM limited_sessions
        )

        SELECT json_build_object(
            'info', json_build_object(
                'id', si.id,
                'name', si.fullname,
                'department', si.department,
                'shift', si.shift,
                'semester', si.semester,
                'section', si.section,
                'total_classes', cc.total_classes,
                'present', cc.present,
                'absent', (cc.total_classes - cc.present)
            ),
            'sessions', ls.sessions
        )
        FROM student_info si
        JOIN class_counts cc ON true
        JOIN last_sessions ls ON true
    `

	tx, err := d.pool.BeginTx(ctx, pgx.TxOptions{
		AccessMode: pgx.ReadOnly,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		// Ends Transaction
		if err == nil {
			_ = tx.Rollback(ctx)
			return
		}
		_ = tx.Commit(ctx)
	}()

	var jsonBlob json.RawMessage

	// This is the correct way for a query that returns ONE JSON blob
	err = tx.QueryRow(ctx, query, pgx.NamedArgs{"student_id": studentId}).Scan(&jsonBlob)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("student with id %d not found", studentId)
		}
		return nil, fmt.Errorf("failed to query attendance overview: %w", err)
	}

	var overview dto.StudentAttendanceOverview
	if err = json.Unmarshal(jsonBlob, &overview); err != nil {
		return nil, fmt.Errorf("failed to collect data from json err:%w", err)
	}

	return &overview, nil
}

func (d *DB) StudentsAttendanceOverview(ctx context.Context, paginate *dto.PaginationDto) (*model.PaginatedResponse[dto.StudentsOverview], error) {

	filterClauses := []string{}
	joinFilterClauses := []string{}

	args := pgx.NamedArgs{
		"limit": paginate.Limit,
		"page":  paginate.Page,
	}

	for key, value := range paginate.Filter {
		filterClauses = append(filterClauses, fmt.Sprintf("%s = @%s", key, key))
		joinFilterClauses = append(joinFilterClauses, fmt.Sprintf("s.%s = @%s", key, key))
		args[key] = value
	}

	filterString := ""
	joinFilterString := ""
	if len(filterClauses) > 0 {
		filterString = "WHERE " + strings.Join(filterClauses, " AND ")
		joinFilterString = "WHERE " + strings.Join(joinFilterClauses, " AND ")
	}

	query := fmt.Sprintf(`
	WITH session_counts AS (
        SELECT
            department,
            shift,
            semester,
            section,
            COUNT(*) AS session_count
        FROM
            attendance_sessions
		%s
        GROUP BY
            department,
            shift,
            semester,
            section
        ),
        attendance_counts AS (
        SELECT
            student_id,
            COUNT(*) AS attended_count
        FROM
            attendance_records
        WHERE
            student_id IN (
            SELECT
                student_id
            FROM
				students
			%s)
        GROUP BY
            student_id
        ),
        student_data AS (
        SELECT
            s.id,
            s.student_id,
            s.fullname,
            COALESCE(sc.session_count, 0) AS total_sessions,
            COALESCE(ac.attended_count, 0) AS total_attended,
            (
            	COALESCE(ac.attended_count, 0) :: decimal / NULLIF(COALESCE(sc.session_count, 0), 0) * 100
            ) :: INT AS percentage
        FROM
            students s
            LEFT JOIN session_counts sc ON s.department = sc.department
            AND s.shift = sc.shift
            AND s.semester = sc.semester
            AND s.section = sc.section
            LEFT JOIN attendance_counts ac ON s.student_id = ac.student_id
		%s
		),
        count_data AS (
        SELECT
            COUNT(*) AS total_rows
        FROM
            student_data
        ),
        paged_data AS (
        SELECT
            *
        FROM
            student_data
        ORDER BY
            percentage DESC
        LIMIT
            @limit OFFSET (@page - 1) * @limit
        )
        SELECT
        json_build_object(
            'total', cd.total_rows,
            'total_pages', CEIL(cd.total_rows :: decimal / @limit),
            'page', @page,
            'limit', @limit,
            'data', COALESCE(jr.students, '[]' :: json)
        ) AS result
        FROM
        count_data cd
        LEFT JOIN (
            SELECT
            json_agg(paged_data.*) AS students
            FROM
            paged_data
        ) jr ON TRUE`, filterString, filterString, joinFilterString)

	tx, err := d.pool.BeginTx(ctx, pgx.TxOptions{
		AccessMode: pgx.ReadOnly,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		// Ends Transaction
		if err == nil {
			_ = tx.Rollback(ctx)
			return
		}
		_ = tx.Commit(ctx)
	}()

	var jsonBlob json.RawMessage

	// This is the correct way for a query that returns ONE JSON blob
	err = tx.QueryRow(ctx, query, args).Scan(&jsonBlob)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("no data found: %w",err)
		}
		return nil, fmt.Errorf("failed to query students overview: %w", err)
	}

	var overview model.PaginatedResponse[dto.StudentsOverview]
	if err = json.Unmarshal(jsonBlob, &overview); err != nil {
		return nil, fmt.Errorf("failed to collect data from json err:%w", err)
	}

	return &overview, nil
}
