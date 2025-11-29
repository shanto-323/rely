package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/shanto-323/rely/internal/server/handler"
	"github.com/shanto-323/rely/internal/server/middleware"
)

func RegisterV1Routes(r *echo.Group, h *handler.Handlers, m *middleware.Middlewares) {
	studentRoute := r.Group("/overview")
	{
		studentRoute.POST("", h.Overview.GetStudentsOverview)
		studentRoute.GET("/:student_id", h.Overview.GetStudentOverview)
	}

}
