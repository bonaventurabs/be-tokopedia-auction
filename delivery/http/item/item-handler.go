package item

import (
	"net/http"
	"strconv"
	"time"

	"github.com/bonaventurabs/be-tokopedia-auction/model"
	"github.com/labstack/echo"
)

func (h *ItemHttpHandler) GetAllItems(c echo.Context) error {
	ctx := c.Request().Context()
	data, err := h.usecase.GetAllItemsDetail(ctx)
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

func (h *ItemHttpHandler) GetItemById(c echo.Context) error {
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

	data, err := h.usecase.GetItemDetailById(ctx, id)
	if err != nil && err.Error() == "record not found" {
		return c.JSON(http.StatusNotFound, model.Response{
			Code:    http.StatusNotFound,
			Status:  http.StatusText(http.StatusNotFound),
			Message: err.Error(),
		})
	} else if err != nil {
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

func (h *ItemHttpHandler) PostItem(c echo.Context) error {
	ctx := c.Request().Context()
	var item model.ItemPost

	if err := c.Bind(&item); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
	}

	// validate start auction and end auction in rfc3339
	start_auct, _ := time.Parse(time.RFC3339, item.StartAuction)
	end_auct, _ := time.Parse(time.RFC3339, item.EndAuction)

	if start_auct.Before(time.Now()) {
		return c.JSON(http.StatusBadRequest, model.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "start auction must be after now",
		})
	}

	if start_auct.After(end_auct) {
		return c.JSON(http.StatusBadRequest, model.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "start auction must be before end auction",
		})
	}

	if err := h.usecase.PostItem(ctx, model.Item{
		Name:         item.Name,
		Description:  item.Description,
		IsAuction:    true,
		IsActive:     true,
		StartAuction: item.StartAuction,
		EndAuction:   item.EndAuction,
		StartPrice:   item.StartPrice,
		SellerId:     item.SellerId,
	}); err != nil {
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
	})
}
