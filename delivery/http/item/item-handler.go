package item

import (
	"net/http"

	"github.com/bonaventurabs/be-tokopedia-auction/model"
	"github.com/labstack/echo"
)

func (h *ItemHttpHandler) GetAllItems(c echo.Context) error {
	ctx := c.Request().Context()
	data, err := h.usecase.GetAllItemsDetail(ctx)
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
		Message: "Success",
		Data:    data,
	})
}
