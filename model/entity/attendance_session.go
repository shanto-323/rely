package entity

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/shanto-323/rely/model"
)

type AttendanceSession struct {
	model.Base
	TeacherID   int              `db:"teacher_id" validate:"required"`
	SubjectCode int              `db:"subject_code" validate:"required"`
	Department  model.Department `db:"department" validate:"required"`
	Shift       model.Shift      `db:"shift" validate:"required,oneof=1 2"`
	Semester    model.Semester   `db:"semester" validate:"required,min=1,max=8"`
	Section     model.Section    `db:"section" validate:"required,oneof=A B C"`
	Valid       bool             `db:"valid"`
}

func (as *AttendanceSession) Validate() error {
	return validator.New().Struct(as)
}

func (as *AttendanceSession) IsValid() error {
	if as.Valid {
		return fmt.Errorf("this session session_id: %s is not valid", as.ID)
	}
	return nil
}
