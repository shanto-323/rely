package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/shanto-323/rely/model/entity"
)

func (d *DB) CreateAttendanceSession(ctx context.Context, session *entity.AttendanceSession, studentIDs []int) (*entity.AttendanceSession, error) {
	if err := session.Validate(); err != nil {
		return nil, fmt.Errorf("invalid attendance session: %w", err)
	}

	query := `
	INSERT INTO attendance_sessions (
		teacher_id,
		subject_code,
		department,
		shift,
		semester,
		section,
		valid,
	)
	VALUES (
		@teacher_id,
		@subject_code,
		@department,
		@shift,
		@semester,
		@section,
		@valid,
	)
	RETURNING *
	`

	// Starts Transaction
	tx, err := d.pool.BeginTx(ctx, pgx.TxOptions{BeginQuery: ""})
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

	rows, err := tx.Query(ctx, query, pgx.NamedArgs{
		"teacher_id":   session.TeacherID,
		"subject_code": session.SubjectCode,
		"department":   session.Department,
		"shift":        session.Shift,
		"semester":     session.Semester,
		"section":      session.Section,
		"valid":        session.Valid,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to insert session: %w", err)
	}

	sess, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[entity.AttendanceSession])
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve session_id: %w", err)
	}

	err = d.CreateAttendanceRecords(ctx, tx, sess.ID, studentIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to create attendance records: %w", err)
	}

	return &sess, nil
}
