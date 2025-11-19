package dto

import "github.com/shanto-323/rely/model"

type AttendanceDto struct {
	TeacherID  int              `json:"teacher_id" validate:"required"`
	SubjectID  int              `json:"subject" validate:"required"`
	Department model.Department `json:"department" validate:"required"`
	Shift      model.Shift      `json:"shift" validate:"required"`
	Semester   model.Semester   `json:"semester" validate:"required"`
	Section    model.Section    `json:"section" validate:"required"`
	StudentIDs []int            `json:"student_ids" validate:"required"`
}

