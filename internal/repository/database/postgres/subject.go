package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/shanto-323/rely/model/entity"
)

func (d *DB) CreateSubject(ctx context.Context, payload *entity.Subject) (*entity.Subject, error) {
	if err := payload.Validate(); err != nil {
		return nil, fmt.Errorf("error validating payload")
	}

	query := `
	INSERT INTO 
	subjects (
		code,
		name,
		credits,
		semester,
	)
	VALUES (
		@code,
		@name,
		@credits,
		@semester,
	)
	RETURNING 
		*
	`

	rows, err := d.pool.Query(ctx, query, pgx.NamedArgs{
		"code":       payload.Code,
		"name":       payload.Name,
		"credits":    payload.Credits,
		"semester":   payload.Semester,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create subject: %w", err)
	}

	subject, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[entity.Subject])
	if err != nil {
		return nil, fmt.Errorf("failed to collect subject data")
	}

	return &subject, nil
}

func (d *DB) GetSubjectByID(ctx context.Context, id uuid.UUID) (*entity.Subject, error) {
	query := `
	SELECT
		*
	FROM 
		subjects
	WHERE
		id = @id
	`

	rows, err := d.pool.Query(ctx, query, pgx.NamedArgs{"id": id})
	if err != nil {
		return nil, fmt.Errorf("failed to select subject: %w", err)
	}

	subject, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[entity.Subject])
	if err != nil {
		return nil, fmt.Errorf("subject not found")
	}

	return &subject, nil
}

func (d *DB) DeleteSubjectByID(ctx context.Context, id uuid.UUID) error {
	query := `
	DELETE FROM 
		subjects
	WHERE 
		id = @id
	`

	result, err := d.pool.Exec(ctx, query, pgx.NamedArgs{"id": id})
	if err != nil {
		return fmt.Errorf("failed to delete subject: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("subject not found")
	}

	return nil
}
