package user

import (
	"context"

	"github.com/bonaventurabs/be-tokopedia-auction/model"
)

type User interface {
	GetUserById(ctx context.Context, id int64) (model.User, error)
}
