package entity

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/shanto-323/rely/model"
)

type User struct {
	ID       uuid.UUID      `db:"id" validate:"required,uuid"`
	UserId   uuid.UUID      `db:"user_id" validate:"required,uuid"` // = ID, Foreign key
	UserType model.UserType `db:"user_type" validate:"required,oneof=SUP MOD TEACHER STUDENT"`
	Token    string         `db:"token" validate:"required,min=10"`
	Blocked  bool           `db:"blocked"`
}

func (u *User) Validate() error {
	return validator.New().Struct(u)
}

func (u *User) IsBlocked() error {
	if u.Blocked {
		return fmt.Errorf("the user with user_id: %d && user_type: %s is currently blocked by system", u.UserId, u.UserType)
	}
	return nil
}
