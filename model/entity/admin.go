package entity

import (
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/shanto-323/rely/model"
)

type Admin struct {
	model.Base
	IssuedBy uuid.UUID `db:"issued_by" validate:"required,uuid"` // UserId
	CaimedBy uuid.UUID `db:"used_by" validate:"required,uuid"`   // UserId
}

func (a *Admin) Validate() error {
	return validator.New().Struct(a)
}
