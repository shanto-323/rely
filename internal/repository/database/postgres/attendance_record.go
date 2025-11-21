package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (d *DB) CreateAttendanceRecords(ctx context.Context, tx pgx.Tx, session_id uuid.UUID, studentIds []int) error {
	rows := make([][]any, len(studentIds))
	for i, r := range studentIds {
		rows[i] = []any{session_id, r}
	}

	_, err := tx.CopyFrom(ctx,
		pgx.Identifier{"attendance_records"},
		[]string{"session_id", "student_id"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return fmt.Errorf("failed to insert attendance batch: %w", err)
	}

	return nil
}



