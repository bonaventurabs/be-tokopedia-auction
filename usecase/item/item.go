package item

import (
	"context"

	"github.com/bonaventurabs/be-tokopedia-auction/model"
)

func (r *Repo) GetAllItemsDetail(ctx context.Context) ([]model.ItemDetail, error) {
	var (
		data []model.ItemDetail
	)

	items, err := r.Item.GetAllItems(ctx)
	if err != nil {
		return model.ItemDetail{}, err
	}

	for _, item := range data {

	}

	return data, nil
}
