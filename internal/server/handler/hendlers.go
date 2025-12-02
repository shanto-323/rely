package handler

import (
	"github.com/shanto-323/rely/internal/server"
	"github.com/shanto-323/rely/internal/service"
)

type Handlers struct {
	Health   *HealthHandler
	Overview *OverviewHandler
	Student *StudentHandler
}

func New(s *server.Server, services *service.Services) *Handlers {
	return &Handlers{
		Health:   NewHealthHandler(s),
		Overview: NewOverviewHandler(services.OverviewService),
		Student: NewStudentHandler(services.Student),
	}
}
