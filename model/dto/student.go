package dto

import (
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type StudentsQueryRequest struct {
	AcademicContext
	Pagination
}

func (r *StudentsQueryRequest) Validate() error {
	if r.Page == nil {
		defaultPage := 1
		r.Page = &defaultPage
	}

	if r.Limit == nil {
		defaultLimit := 10
		r.Limit = &defaultLimit
	}

	validator := validator.New()
	return validator.Struct(r)
}

type StudentIDRequest struct {
	StudentID int `param:"student_id" validate:"required"`
}

func (r *StudentIDRequest) Validate() error {
	validator := validator.New()
	return validator.Struct(r)
}

type CreateStudentRequest struct {
	StudentID    int    `json:"student_id" validate:"required,min=6"`
	FullName     string `json:"fullname" validate:"required,min=3,max=100"`
	Email        string `json:"email" validate:"required,email"`
	Phone        string `json:"phone" validate:"required,e164"`
	Registration int    `json:"registration" validate:"required,min=11"`
	Department   string `json:"department" validate:"required"`
	Shift        string `json:"shift" validate:"required"`
	Semester     string `json:"semester" validate:"required,min=1,max=8"`
	Section      string `json:"section" validate:"required"`
}

func (r *CreateStudentRequest) Validate() error {
	validator := validator.New()
	return validator.Struct(r)
}

type CreateStudentResponse struct {
	Id         uuid.UUID `json:"id"`
	StudentID  int       `json:"student_id"`
	FullName   string    `json:"fullname"`
	Department string    `json:"department"`
	Semester   string    `json:"semester"`
	Section    string    `json:"section"`
}
