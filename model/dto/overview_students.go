package dto

import "github.com/go-playground/validator"

type OverviewStudentsQueryRequest struct {
	AcademicContext
	Pagination
}

func (r *OverviewStudentsQueryRequest) Validate() error {
	validator := validator.New()
	return validator.Struct(r)
}

type StudentsOverview struct {
	StudentID     int    `json:"student_id"`
	Fullname      string `json:"fullname"`
	TotalSessions int    `json:"total_sessions"`
	TotalAttended int    `json:"total_attended"`
	Percentage    int    `json:"percentage"`
}
