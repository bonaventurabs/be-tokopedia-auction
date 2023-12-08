package user

import (
	"net/http"
	"strconv"

	"github.com/bonaventurabs/be-tokopedia-auction/model"
	"github.com/labstack/echo"
)

func (h *UserHttpHandler) GetUserById(c echo.Context) error {
	ctx := c.Request().Context()
	rId := c.Param("id")

	if rId == "" {
		return c.JSON(http.StatusBadRequest, model.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "id is mandatory",
		})
	}

	id, err := strconv.ParseInt(rId, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "invalid id",
		})
	}

	data, err := h.usecase.GetUserById(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.Response{
			Code:    http.StatusNotFound,
			Status:  http.StatusText(http.StatusNotFound),
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
