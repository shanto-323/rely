package handler

import (
	"net/http"
	"strconv"

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
	return HandleIdBasedPath(
		func(c echo.Context) (any, error) {
			studentID := c.Param("student_id")

			id, err := strconv.Atoi(studentID)
			if err != nil {
				return nil, err
			}
			return h.service.StudentAttendanceOverview(c, id)
		}, JSONResponseHandler{status: http.StatusOK})(c)
}

func (h *OverviewHandler) GetStudentsOverview(c echo.Context) error {
	return Handle(
		func(c echo.Context, req *dto.PaginationDto) (*model.PaginatedResponse[dto.StudentsOverview], error) {
			return h.service.StudentsAttendanceOverview(c, req)
		},
		http.StatusOK,
		&dto.PaginationDto{},
	)(c)
}
