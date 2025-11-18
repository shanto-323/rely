package entity

import (
	"time"

	"github.com/go-playground/validator"
	"github.com/shanto-323/rely/model"
)

type AttendanceSession struct {
	model.Base
	TeacherID  int              `db:"teacher_id" validate:"required"`
	SubjectID  int              `db:"subject" validate:"required"`
	ClassDate  time.Time        `db:"class_date" validate:"required"`
	Department model.Department `db:"department" validate:"required"`
	Shift      model.Shift      `db:"shift" validate:"required"`
	Semester   model.Semester   `db:"semester" validate:"required,min=1,max=8"`
	Section    model.Section    `db:"section" validate:"required,oneof=A B C"`
	IsValid    bool             `db:"valid"`
}

func (as *AttendanceSession) Validate() error {
	return validator.New().Struct(as)
}
