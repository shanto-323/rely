package dto

import "github.com/shanto-323/rely/model"

// DTO for receiving data from client
type StudentRequestDTO struct {
	StudentID    int    `json:"student_id" validate:"required,min=6"`
	FullName     string `json:"fullname" validate:"required,min=3,max=100"`
	Email        string `json:"email" validate:"required,email"`
	Phone        string `json:"phone" validate:"required,e164"`
	Registration int    `json:"registration" validate:"required,min=11"`
	Department   string `json:"department" validate:"required"` // string representation
	Shift        string `json:"shift" validate:"required"`      // string representation
	Semester     int    `json:"semester" validate:"required,min=1,max=8"`
	Section      string `json:"section" validate:"required"` // string representation
}

// DTO for sending data to client
type StudentResponseDTO struct {
	model.Base
	StudentID    int    `json:"student_id"`
	FullName     string `json:"fullname"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Registration int    `json:"registration"`
	Department   string `json:"department"`
	Shift        string `json:"shift"`
	Semester     int    `json:"semester"`
	Section      string `json:"section"`
}

// Public DTO for Student (for non-admin / client-facing endpoints)
type StudentPublicDTO struct {
	StudentID  int    `json:"student_id"`
	FullName   string `json:"fullname"`
	Department string `json:"department"`
	Semester   int    `json:"semester"`
	Section    string `json:"section"`
}

