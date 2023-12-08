package user

import (
	"github.com/labstack/echo"
)

type UserHttpHandler struct {
	usecase Usecases
}

func NewUserHTTP(e *echo.Group, usecase Usecases) {
	handler := &UserHttpHandler{
		usecase: usecase,
	}
	e.GET("/users/:id", handler.GetUserById)
}
