package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shanto-323/rely/internal/service"
	"github.com/shanto-323/rely/model"
	"github.com/shanto-323/rely/model/dto"
)

type OverviewHandler struct {
	service *service.OverviewService
}

func NewOverviewHandler(service *service.OverviewService) *OverviewHandler {
	return &OverviewHandler{
		service: service,
	}
}

func (h *OverviewHandler) GetStudentOverview(c echo.Context) error {
	return Handle(
		func(c echo.Context, studentId *dto.StudentIDRequest) (*dto.StudentAttendanceOverview, error) {
			return h.service.StudentAttendanceOverview(c, studentId.StudentID)
		},
		http.StatusOK,
		&dto.StudentIDRequest{},
	)(c)
}

func (h *OverviewHandler) GetStudentsOverview(c echo.Context) error {
	return Handle(
		// need to re-build
		func(c echo.Context, req *dto.OverviewStudentsQueryRequest) (*model.PaginatedResponse[dto.StudentsOverview], error) {
			return h.service.StudentsAttendanceOverview(c, req)
		},
		http.StatusOK,
		&dto.OverviewStudentsQueryRequest{},
	)(c)
}
