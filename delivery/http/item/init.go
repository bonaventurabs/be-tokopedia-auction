package item

import (
	"github.com/labstack/echo"
)

type ItemHttpHandler struct {
	usecase Usecases
}

func NewItemHTTP(e *echo.Group, usecase Usecases) {
	handler := &ItemHttpHandler{
		usecase: usecase,
	}
	e.GET("/items", handler.GetAllItems)
}
