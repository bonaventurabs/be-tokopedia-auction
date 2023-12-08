package health

import (
	"github.com/labstack/echo"
)

type HealthHttpHandler struct {
	usecase Usecases
}

func NewHealthHTTP(e *echo.Echo, usecase Usecases) {
	handler := &HealthHttpHandler{
		usecase: usecase,
	}
	e.GET("/health", handler.HandlerHealthCheck)
}
