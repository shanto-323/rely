package entity

import (
	"github.com/go-playground/validator"
	"github.com/shanto-323/rely/model"
)

type Student struct {
	model.Base
	StudentID    int              `db:"student_id" validate:"required,min=6"` // = Roll
	FullName     string           `db:"fullname" validate:"required,min=3,max=100"`
	Email        string           `db:"email" validate:"required,email"`
	Phone        string           `db:"phone" validate:"required,e164"`
	Registration int              `db:"registration" validate:"required,min=11"`
	Department   model.Department `db:"department" validate:"required"`
	Shift        model.Shift      `db:"shift" validate:"required"`
	Semester     model.Semester   `db:"semester" validate:"required,min=1,max=8"`
	Section      model.Section    `db:"section" validate:"required"`
}

func (s *Student) Validate() error {
	return validator.New().Struct(s)
}

