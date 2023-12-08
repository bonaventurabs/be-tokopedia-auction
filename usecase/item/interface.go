package item

import (
	"context"

	"github.com/bonaventurabs/be-tokopedia-auction/model"
)

type Item interface {
	GetAllItems(ctx context.Context) ([]model.Item, error)
	GetImagesByItemId(ctx context.Context, id int64) ([]model.ItemImage, error)
	GetSellerById(ctx context.Context, id int64) (model.Seller, error)
	GetBidByItemId(ctx context.Context, id int64) ([]model.Bid, error)
}
