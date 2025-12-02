package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shanto-323/rely/internal/service"
	"github.com/shanto-323/rely/model"
	"github.com/shanto-323/rely/model/dto"
	"github.com/shanto-323/rely/model/entity"
)

type StudentHandler struct {
	service *service.StudentService
}

func NewStudentHandler(service *service.StudentService) *StudentHandler {
	return &StudentHandler{
		service: service,
	}
}

func (s *StudentHandler) GetStudents(c echo.Context) error {
	return Handle(
		func(c echo.Context, query *dto.StudentsQueryRequest) (*model.PaginatedResponse[entity.Student], error) {
			return s.service.GetStudents(c, query)
		},
		http.StatusOK,
		&dto.StudentsQueryRequest{},
	)(c)
}
