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

func (r *Repo) GetItemDetailById(ctx context.Context, id int64) (model.ItemDetail, error) {
	var (
		data model.ItemDetail
	)

	item, err := r.Item.GetItemById(ctx, id)
	if err != nil {
		return model.ItemDetail{}, err
	}

	data.Item = item
	data.Seller, _ = r.Item.GetSellerById(ctx, item.SellerId)
	data.Images, _ = r.Item.GetImagesByItemId(ctx, item.Id)

	bid, _ := r.Item.GetBidByItemId(ctx, item.Id)
	data.Bid = model.BidDetail{
		TotalBid:   int64(len(bid)),
		HighestBid: bid[len(bid)-1].BidPrice,
		HistoryBid: bid,
	}

	return data, nil
}

func (r *Repo) PostItem(ctx context.Context, item model.Item) error {
	const placeholder = "https://ucarecdn.com/7f4d7270-607d-4cfa-81cd-a3c28ca889e7/-/preview/500x500/-/quality/smart/-/format/auto/"

	err := r.Item.PostItem(ctx, item)
	if err != nil {
		return err
	}

	err = r.Item.PostItemImages(ctx, []model.ItemImage{
		{
			Image:  placeholder,
			ItemId: item.Id,
		},
	})
	if err != nil {
		return err
	}

	return nil
}
