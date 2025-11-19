package entity

import (
	"fmt"

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
	Valid     bool      `db:"valid"`
}

func (t *Token) Validate() error {
	return validator.New().Struct(t)
}

func (t *Token) IsValid() error {
	if t.Valid {
		return fmt.Errorf("this token token_id: %s is not valid", t.ID)
	}
	return nil
}

