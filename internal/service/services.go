package service

import (
	"github.com/shanto-323/rely/internal/server"
)

type Services struct {
	OverviewService *OverviewService
}

func New(s *server.Server) *Services {
	return &Services{
		OverviewService: NewOverviewService(s.Repository.DatabaseDriver),
	}
}
