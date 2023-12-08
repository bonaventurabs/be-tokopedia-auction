package bid

import (
	"context"

	"github.com/bonaventurabs/be-tokopedia-auction/model"
)

type Usecases interface {
	PostBid(ctx context.Context, bid model.BidPost) (model.Bid, error)
}
