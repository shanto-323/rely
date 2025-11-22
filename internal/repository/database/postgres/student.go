package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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
