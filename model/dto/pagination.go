package dto

import "github.com/go-playground/validator"

type PaginationDto struct {
	Limit  int            `json:"limit" validate:"required,gte=10,lte=100"`
	Page   int            `json:"page" validate:"required,gte=1"`
	Filter map[string]any `json:"filter" validate:"omitempty"`
}

func (p *PaginationDto) Validate() error {
	// bug :filter should be one of department,shift,semester, section,

	v := validator.New()
	return v.Struct(p)
}
