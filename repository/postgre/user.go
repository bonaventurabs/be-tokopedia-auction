package postgre

import (
	"context"

	"github.com/bonaventurabs/be-tokopedia-auction/model"
)

func (db *DB) GetUserById(ctx context.Context, id int64) (model.User, error) {
	var (
		item model.User
	)

	query := `SELECT id, created_at, updated_at, username, name, email, image 
		FROM users WHERE id = $1`
	err := db.DB.QueryRowContext(ctx, query, id).Scan(
		&item.Id, &item.CreatedAt, &item.UpdatedAt, &item.Username, &item.Name, &item.Email, &item.Image)
	if err != nil {
		return model.User{}, err
	}

	return item, nil
}
