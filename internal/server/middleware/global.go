package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shanto-323/rely/internal/server"
)

type Global struct {
	s *server.Server
}

func NewGlobal(s *server.Server) *Global {
	return &Global{
		s: s,
	}
}

func (g *Global) CROS() echo.MiddlewareFunc {
    return middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"*"},
        AllowMethods: []string{
            echo.GET,
            echo.POST,
            echo.PUT,
            echo.DELETE,
            echo.OPTIONS,
        },
        AllowHeaders: []string{
            echo.HeaderOrigin,
            echo.HeaderContentType,
            echo.HeaderAccept,
            echo.HeaderAuthorization,
        },
        AllowCredentials: true,
    })
}

