package dto

import (
	"time"

	"github.com/google/uuid"
)

type StudentAttendanceOverview struct {
	Info     Info      `json:"info"`
	Sessions []Session `json:"sessions"`
}

type Info struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Department   string    `json:"department"`
	Shift        string    `json:"shift"`
	Semester     string    `json:"semester"`
	Section      string    `json:"section"`
	TotalClasses int       `json:"total_classes"`
	Present      int       `json:"present"`
	Absent       int       `json:"absent"`
}

type Session struct {
	SessionID   uuid.UUID `json:"session_id"`
	Teacher     Teacher   `json:"teacher"`
	SubjectCode int       `json:"subject_code"`
	CreatedAt   time.Time `json:"created_at"`
	Present     bool      `json:"present"`
}

type Teacher struct {
	ID       uuid.UUID `json:"id"`
	Fullname string    `json:"fullname"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
}
