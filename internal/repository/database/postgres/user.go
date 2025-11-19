package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/shanto-323/rely/model/entity"
)

func (d *DB) CreateUser(ctx context.Context, tx pgx.Tx, payload *entity.User) (*entity.User, error) {
	if err := payload.Validate(); err != nil {
		return nil, fmt.Errorf("error validating payload")
	}

	query := `
	INSERT INTO users (
		id,
		user_id,
		user_type,
		token,
		blocked
	)
	VALUES (
		@id,
		@user_id,
		@user_type,
		@token,
		@blocked
	)
	RETURNING *
	`

	rows, err := tx.Query(ctx, query, pgx.NamedArgs{
		"id":        payload.ID,
		"user_id":   payload.UserId,
		"user_type": payload.UserType,
		"token":     payload.Token,
		"blocked":   payload.Blocked,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to execute user query: %w", err)
	}

	user, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[entity.User])
	if err != nil {
		return nil, fmt.Errorf("failed to map database row to user struct: %w", err)
	}

	return &user, nil
}

func (d *DB) GetUserByID(ctx context.Context, userId uuid.UUID) (*entity.User, error) {
	query := `
	SELECT 
		*
	FROM 
		users
	WHERE 
		user_id=@user_id
	`

	rows, err := d.pool.Query(ctx, query, &pgx.NamedArgs{
		"user_id": userId,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to execute user query: %w", err)
	}

	user, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[entity.User])
	if err != nil {
		return nil, fmt.Errorf("failed to map database row to user struct: %w", err)
	}

	return &user, nil
}
