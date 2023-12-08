package item

import (
	"context"

	"github.com/bonaventurabs/be-tokopedia-auction/model"
)

type Usecases interface {
	GetAllItemsDetail(ctx context.Context) (model.ItemDetailPagination, error)
	GetItemDetailById(ctx context.Context, id int64) (model.ItemDetail, error)
	PostItem(ctx context.Context, item model.Item) error
}
