package postgre

import (
	"context"

	"github.com/bonaventurabs/be-tokopedia-auction/model"
)

func (db *DB) GetSellerById(ctx context.Context, id int64) (model.Seller, error) {
	var (
		item model.Seller
	)

	query := `SELECT 
			id, 
			created_at, 
			updated_at, 
			username, 
			name, 
			email, 
			image 
		FROM seller WHERE id = $1`
	err := db.DB.QueryRowContext(ctx, query, id).Scan(
		&item.Id,
		&item.CreatedAt,
		&item.UpdatedAt,
		&item.Username,
		&item.Name,
		&item.Email,
		&item.Image,
	)
	if err != nil {
		return model.Seller{}, err
	}

	return item, nil
}
