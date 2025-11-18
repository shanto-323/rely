package handler

import (
	"github.com/shanto-323/rely/internal/server"
	"github.com/shanto-323/rely/internal/service"
)

type Handlers struct {
	HealthHandler  *HealthHandler
}

func New(s *server.Server, sr *service.Services) *Handlers {
	return &Handlers{
		HealthHandler:  NewHealthHandler(s),
	}
}
