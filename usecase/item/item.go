package item

import (
	"context"

	"github.com/bonaventurabs/be-tokopedia-auction/model"
)

func (r *Repo) GetAllItemsDetail(ctx context.Context) (model.ItemDetailPagination, error) {
	var (
		data model.ItemDetailPagination
	)
	data.Page = 1
	data.TotalCount = 1

	items, err := r.Item.GetAllItems(ctx)
	if err != nil {
		return model.ItemDetailPagination{}, err
	}

	data.PerPage = int64(len(items))
	data.TotalCount = int64(len(items))

	for _, item := range items {
		var itemDetail model.ItemDetail
		itemDetail.Item = item
		itemDetail.Seller, _ = r.Item.GetSellerById(ctx, item.SellerId)
		itemDetail.Images, _ = r.Item.GetImagesByItemId(ctx, item.Id)

		bid, _ := r.Item.GetBidByItemId(ctx, item.Id)
		itemDetail.Bid = model.BidDetail{
			TotalBid:   int64(len(bid)),
			HighestBid: bid[len(bid)-1].BidPrice,
			HistoryBid: bid,
		}
		data.Records = append(data.Records, itemDetail)
	}

	return data, nil
}
