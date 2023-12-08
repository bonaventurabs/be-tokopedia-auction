package user

import (
	"context"

	"github.com/bonaventurabs/be-tokopedia-auction/model"
)

func (r *Repo) GetUserById(ctx context.Context, id int64) (model.User, error) {
	data, err := r.User.GetUserById(ctx, id)
	if err != nil {
		return model.User{}, err
	}

	return data, nil
}
