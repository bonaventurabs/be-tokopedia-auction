package bid

import (
	"github.com/labstack/echo"
)

type BidHttpHandler struct {
	usecase Usecases
}

func NewBidHTTP(e *echo.Group, usecase Usecases) {
	handler := &BidHttpHandler{
		usecase: usecase,
	}
	e.POST("/bid", handler.PostBid)
}
