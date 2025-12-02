package entity

import (
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/shanto-323/rely/model"
)

type AttendanceSession struct {
	model.Base
	TeacherID   uuid.UUID `db:"teacher_id" validate:"required"`
	SubjectCode int       `db:"subject_code" validate:"required"`
	Department  string    `db:"department" validate:"required"`
	Shift       string    `db:"shift" validate:"required,oneof=1 2"`
	Semester    string    `db:"semester" validate:"required,min=1,max=8"`
	Section     string    `db:"section" validate:"required,oneof=A B C"`
}

func (as *AttendanceSession) Validate() error {
	return validator.New().Struct(as)
}
