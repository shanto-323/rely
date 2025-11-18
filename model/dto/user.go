package dto

// DTO for receiving data from client
type UserRequestDTO struct {
	UserId int    `json:"user_id" validate:"required"`
	Token  string `json:"token" validate:"required"`
}
