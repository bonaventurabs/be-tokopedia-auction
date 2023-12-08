package bid

import (
	"context"

	"github.com/bonaventurabs/be-tokopedia-auction/model"
)

type Bid interface {
	PostBid(ctx context.Context, bid model.BidPost) (model.Bid, error)
	GetBidByItemId(ctx context.Context, id int64) ([]model.Bid, error)
	GetHighestBidByItemId(ctx context.Context, id int64) (model.Bid, error)
}
