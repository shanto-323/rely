package handler

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shanto-323/rely/internal/server/middleware"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type HandleIdBasedFunc func(c echo.Context) (any, error)

func handleIdBasedRequest(
	c echo.Context,
	handler func(c echo.Context) (any, error),
	responseHandler ResponseHandler,
) error {
	start := time.Now()
	method := c.Request().Method
	path := c.Path()

	span := trace.SpanFromContext(c.Request().Context())

	// Get context-enhanced logger
	loggerBuilder := middleware.GetLogger(c).With().
		Str("operation", responseHandler.GetOperation()).
		Str("method", method).
		Str("path", path)

	logger := loggerBuilder.Logger()

	handlerStart := time.Now()
	result, err := handler(c)
	handlerDuration := time.Since(handlerStart)

	if err != nil {
		totalDuration := time.Since(start)
		logger.Error().
			Err(err).
			Dur("handler_duration", handlerDuration).
			Dur("total_duration", totalDuration).
			Msg("handler execution failed")
		span.RecordError(err)
		span.SetAttributes(
			attribute.String("handler.status", "failed"),
			attribute.Int64("handler.duration_ms", handlerDuration.Milliseconds()),
		)
		return err
	}

	totalDuration := time.Since(start)
	span.SetAttributes(
		attribute.String("handler.status", "success"),
		attribute.Int64("handler.duration_ms", handlerDuration.Milliseconds()),
		attribute.Int64("total.duration_ms", totalDuration.Milliseconds()),
	)

	logger.Info().
		Dur("handler_duration", handlerDuration).
		Dur("total_duration", totalDuration).
		Msg("request completed successfully")

	return responseHandler.Handle(c, result)
}

func HandleIdBasedPath(
	handler HandleIdBasedFunc,
	responseHandler ResponseHandler,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		return handleIdBasedRequest(
			c,
			handler,
			responseHandler,
		)
	}
}
