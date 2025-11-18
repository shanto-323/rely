package entity

import (
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/shanto-323/rely/model"
)

type Token struct {
	model.Base
	Code      string    `db:"token" validate:"required"`
	IssuedBy  uuid.UUID `db:"issued_by" validate:"required,uuid"` // UserId
	CaimedBy  uuid.UUID `db:"used_by" validate:"required,uuid"`   // UserId
	IsClaimed bool      `db:"claimed"`
	IsValid   bool      `db:"valid"`
}

func (t *Token) Validate() error {
	return validator.New().Struct(t)
}

