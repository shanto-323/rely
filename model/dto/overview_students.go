package dto

type StudentsOverview struct {
	StudentID     int    `json:"student_id"`
	Fullname      string `json:"fullname"`
	TotalSessions int    `json:"total_sessions"`
	TotalAttended int    `json:"total_attended"`
	Percentage    int    `json:"percentage"`
}
