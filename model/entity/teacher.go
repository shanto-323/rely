package entity

import (
	"github.com/go-playground/validator"
	"github.com/shanto-323/rely/model"
)

type Teacher struct {
	model.Base
	TeacherID int    `db:"teacher_id" validate:"required,min=1"`
	FullName  string `db:"fullname" validate:"required,min=3,max=100"`
	Email     string `db:"email" validate:"required,email"`
	Phone     string `db:"phone" validate:"required,e164"`
}

func (t *Teacher) Validate() error {
	return validator.New().Struct(t)
}
