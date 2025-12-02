package postgres

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/shanto-323/rely/model"
	"github.com/shanto-323/rely/model/entity"
)

func (d *DB) CreateStudent(ctx context.Context, payload *entity.Student, userPayload *entity.User) (*entity.Student, *entity.User, error) {
	if err := payload.Validate(); err != nil {
		return nil, nil, fmt.Errorf("error validating payload")
	}

	query := `
	INSERT INTO
		students (
			student_id,
			fullname,
			email,
			phone,
			registration,
			department,
			shift,
			semester,
			section,
		)
	VALUES
		(
			@student_id,
			@fullname,
			@email,
			@phone,
			@registration,
			@department,
			@shift,
			@semester,
			@section,
		)
	RETURNING
		*
	`

	// Starts Transaction
	tx, err := d.pool.BeginTx(ctx, pgx.TxOptions{
		AccessMode: pgx.ReadWrite,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		// Ends Transaction
		if err == nil {
			_ = tx.Rollback(ctx)
			return
		}
		_ = tx.Commit(ctx)
	}()

	rows, err := tx.Query(ctx, query, pgx.NamedArgs{
		"student_id":   payload.StudentID,
		"fullname":     payload.FullName,
		"email":        payload.Email,
		"phone":        payload.Phone,
		"registration": payload.Registration,
		"department":   payload.Department,
		"shift":        payload.Shift,
		"semester":     payload.Semester,
		"section":      payload.Section,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to execute create todo query")
	}

	student, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[entity.Student])
	if err != nil {
		return nil, nil, fmt.Errorf("failed to collect data from table")
	}

	// Create user by ID = user.UserId
	userPayload.UserId = student.ID
	user, err := d.CreateUser(ctx, tx, userPayload)
	if err != nil {
		return nil, nil, err
	}

	return &student, user, nil
}

func (d *DB) GetStudentByID(ctx context.Context, id uuid.UUID) (*entity.Student, error) {
	query := `
		SELECT
			*
		FROM
			students
		WHERE
			id = @id
	`

	rows, err := d.pool.Query(ctx, query, pgx.NamedArgs{
		"id": id,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to execute student select query")
	}

	student, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[entity.Student])
	if err != nil {
		return nil, fmt.Errorf("student not found")
	}

	return &student, nil
}

func (d *DB) GetStudentByStudentID(ctx context.Context, studentId int) (*entity.Student, error) {
	query := `
		SELECT
			*
		FROM
			students
		WHERE
			student_id = @student_id
	`

	rows, err := d.pool.Query(ctx, query, pgx.NamedArgs{
		"student_id": studentId,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to execute student select query: %w", err)
	}

	student, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[entity.Student])
	if err != nil {
		return nil, fmt.Errorf("student not found: %w", err)
	}

	return &student, nil
}

func (d *DB) DeleteStudentByID(ctx context.Context, id uuid.UUID) error {
	query := `
        DELETE FROM 
			students
        WHERE 
			id = @id
    `

	_, err := d.pool.Exec(ctx, query, pgx.NamedArgs{
		"id": id,
	})
	if err != nil {
		return fmt.Errorf("failed to delete student: %w", err)
	}

	return nil
}

func (d *DB) GetStudents(ctx context.Context, page, limit int, filter map[string]string) (*model.PaginatedResponse[entity.Student], error) {
	filterClauses := []string{}

	args := pgx.NamedArgs{
		"page":  page,
		"limit": limit,
	}

	for key, value := range filter {
		filterClauses = append(filterClauses, fmt.Sprintf("%s = @%s", key, key))
		args[key] = value
	}

	filterString := ""
	if len(filterClauses) > 0 {
		filterString = "WHERE " + strings.Join(filterClauses, " AND ")
	}

	query := fmt.Sprintf(`
		WITH 
			total_students AS (
		        SELECT COUNT(s.id) AS total
		        FROM students s
		       	%s 
		    ),
		    students AS (
		    	SELECT 
					* 
				FROM 
					students s
				%s
				ORDER BY
					student_id ASC 
				LIMIT 
					%d OFFSET (%d-1) * %d
		    )
		SELECT json_build_object(
		    'data', sub.students_json,
		   	'page',%d,
		   	'limit',%d,
		    'total', total_students.total,
		    'total_page', CEIL(total_students.total::numeric / %d)
		)
		FROM total_students
		CROSS JOIN (
		    SELECT json_agg(s) AS students_json
		    FROM students s
		) sub;
	`, filterString, filterString, limit, page, limit, page, limit, limit)

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

	// This returns ONE JSON blob
	err = tx.QueryRow(ctx, query, args).Scan(&jsonBlob)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("no data found: %w", err)
		}
		return nil, fmt.Errorf("failed to query students overview: %w", err)
	}

	var student model.PaginatedResponse[entity.Student]
	if err = json.Unmarshal(jsonBlob, &student); err != nil {
		return nil, fmt.Errorf("failed to collect data from json err:%w", err)
	}

	return &student, nil
}
