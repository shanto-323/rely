package service

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shanto-323/rely/internal/repository/database"
	"github.com/shanto-323/rely/model"
	"github.com/shanto-323/rely/model/dto"
)

type OverviewService struct {
	db database.Driver
}

func NewOverviewService(db database.Driver) *OverviewService {
	return &OverviewService{
		db: db,
	}
}

func (o *OverviewService) StudentAttendanceOverview(c echo.Context, studentId int) (*dto.StudentAttendanceOverview, error) {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	student, err := o.db.GetStudentByStudentID(ctx, studentId)
	if err != nil {
		return nil, err
	}

	return o.db.StudentAttendanceOverview(ctx, student.ID)
}

func (o *OverviewService) StudentsAttendanceOverview(c echo.Context, filter *dto.PaginationDto) (*model.PaginatedResponse[dto.StudentsOverview], error) {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	return o.db.StudentsAttendanceOverview(ctx, filter)
}
