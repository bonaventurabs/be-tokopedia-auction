package item

import (
	"context"

	"github.com/bonaventurabs/be-tokopedia-auction/model"
)

type Usecases interface {
	GetAllItemsDetail(ctx context.Context) (model.ItemDetailPagination, error)
}
