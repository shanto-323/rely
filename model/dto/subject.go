package dto

type SubjectDTO struct {
	Code     int16  `json:"code" validate:"required,min=5"`
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Credits  int    `json:"credits" validate:"required,min=1,max=4"`
	Semester int    `json:"semester" validate:"required,min=1,max=8"`
}
