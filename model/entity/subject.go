package entity

import (
	"github.com/go-playground/validator"
	"github.com/shanto-323/rely/model"
)

type Subject struct {
	model.Base
	Code     int            `db:"code" validate:"required,min=5"`
	Name     string         `db:"name" validate:"required,min=3,max=100"`
	Credits  int            `db:"credits" validate:"required,min=1,max=4"`
	Semester model.Semester `db:"semester" validate:"required,min=1,max=8"`
}

func (c *Subject) Validate() error {
	return validator.New().Struct(c)
}
