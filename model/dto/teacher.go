package dto

import "github.com/shanto-323/rely/model"

// DTO for receiving data from client (request)
type TeacherRequestDTO struct {
	TeacherID int    `json:"teacher_id" validate:"required,min=1"`
	FullName  string `json:"fullname" validate:"required,min=3,max=100"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"required,e164"`
}

// DTO for sending data to client (admin or internal)
type TeacherResponseDTO struct {
	model.Base
	TeacherID int       `json:"teacher_id"`
	FullName  string    `json:"fullname"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
}

// DTO for sending data to public / non-admin (minimal info)
type TeacherPublicDTO struct {
	TeacherID int    `json:"teacher_id"`
	FullName  string `json:"fullname"`
}

