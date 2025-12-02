package service

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shanto-323/rely/internal/repository/database"
	"github.com/shanto-323/rely/model"
	"github.com/shanto-323/rely/model/dto"
	"github.com/shanto-323/rely/model/entity"
)

type StudentService struct {
	db database.Driver
}

func NewStudentService(db database.Driver) *StudentService {
	return &StudentService{
		db: db,
	}
}

func (s *StudentService) GetStudents(c echo.Context, query *dto.StudentsQueryRequest) (*model.PaginatedResponse[entity.Student], error) {
	page := *query.Page
	limit := *query.Limit
	filter := query.GetFilter()

	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	return s.db.GetStudents(ctx, page, limit, filter)
}
