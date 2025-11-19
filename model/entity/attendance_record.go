package entity

import (
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type AttendanceRecord struct {
	SessionId uuid.UUID `db:"session_id" validate:"required,uuid"`
	StudentID int     `db:"student_id" validate:"required,min=1"`
}

func (ar *AttendanceRecord) Validate() error {
	return validator.New().Struct(ar)
}
