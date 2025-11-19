package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/shanto-323/rely/model/entity"
)

func (d *DB) CreateTeacher(ctx context.Context, payload *entity.Teacher, userPayload *entity.User) (*entity.Teacher, *entity.User, error) {
	if err := payload.Validate(); err != nil {
		return nil, nil, fmt.Errorf("error validating payload: %w", err)
	}

	query := `
	INSERT INTO teachers (
		teacher_id,
		fullname,
		email,
		phone,
	)
	VALUES (
		@teacher_id,
		@fullname,
		@email,
		@phone,
	)
	RETURNING *
	`

	tx, err := d.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
			return
		}
		_ = tx.Commit(ctx)
	}()

	rows, err := tx.Query(ctx, query, pgx.NamedArgs{
		"teacher_id": payload.TeacherID,
		"fullname":   payload.FullName,
		"email":      payload.Email,
		"phone":      payload.Phone,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to insert teacher: %w", err)
	}

	teacher, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[entity.Teacher])
	if err != nil {
		return nil, nil, fmt.Errorf("failed to collect teacher data: %w", err)
	}

	// Create user by ID = user.UserId
	userPayload.UserId = teacher.ID
	user, err := d.CreateUser(ctx, tx, userPayload)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create user for teacher: %w", err)
	}

	return &teacher, user, nil
}

func (d *DB) GetTeacherByID(ctx context.Context, id uuid.UUID) (*entity.Teacher, error) {
	query := `
	SELECT *
	FROM teachers
	WHERE id = @id
	`

	rows, err := d.pool.Query(ctx, query, pgx.NamedArgs{"id": id})
	if err != nil {
		return nil, fmt.Errorf("failed to select teacher: %w", err)
	}

	teacher, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[entity.Teacher])
	if err != nil {
		return nil, fmt.Errorf("teacher not found")
	}

	return &teacher, nil
}

func (d *DB) DeleteTeacherByID(ctx context.Context, id uuid.UUID) error {
	query := `
	DELETE FROM teachers
	WHERE id = @id
	`

	result, err := d.pool.Exec(ctx, query, pgx.NamedArgs{"id": id})
	if err != nil {
		return fmt.Errorf("failed to delete teacher: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("teacher not found")
	}

	return nil
}

