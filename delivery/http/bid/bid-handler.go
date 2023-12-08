package bid

import (
	"net/http"

	"github.com/bonaventurabs/be-tokopedia-auction/model"
	"github.com/labstack/echo"
)

func (h *BidHttpHandler) PostBid(c echo.Context) error {
	ctx := c.Request().Context()
	var u model.BidPost
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
	}
	data, err := h.usecase.PostBid(ctx, u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
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
