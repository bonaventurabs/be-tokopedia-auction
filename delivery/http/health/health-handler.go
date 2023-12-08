package health

import (
	"net/http"

	"github.com/bonaventurabs/be-tokopedia-auction/model"
	"github.com/labstack/echo"
)

func (h *HealthHttpHandler) HandlerHealthCheck(c echo.Context) error {
	ctx := c.Request().Context()
	err := h.usecase.HealthCheck(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Message: "API is running",
	})
}
