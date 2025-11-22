package entity

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/shanto-323/rely/model"
)

type Token struct {
	model.Base
	Token     string    `db:"token" validate:"required"`
	IssuedBy uuid.UUID `db:"issued_by" validate:"required,uuid"` // UserId
	ClaimedBy uuid.UUID `db:"claimed_by" validate:"required,uuid"`   // UserId
	Valid    bool      `db:"valid"`
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
