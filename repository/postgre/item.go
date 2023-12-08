package postgre

import (
	"context"

	"github.com/bonaventurabs/be-tokopedia-auction/model"
)

func (db *DB) GetAllItems(ctx context.Context) ([]model.Item, error) {
	var (
		items []model.Item
	)

	query := `SELECT 
			id, 
			created_at, 
			updated_at, 
			name, 
			description, 
			is_auction, 
			is_active,
			start_auct,
			end_auct,
			start_price,
			seller_id 
		FROM item
		`
	rows, err := db.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item model.Item
		err := rows.Scan(
			&item.Id,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.Name,
			&item.Description,
			&item.IsAuction,
			&item.IsActive,
			&item.StartAuction,
			&item.EndAuction,
			&item.StartPrice,
			&item.SellerId,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (db *DB) GetImagesByItemId(ctx context.Context, id int64) ([]model.ItemImage, error) {
	var (
		items []model.ItemImage
	)

	query := `SELECT 
			id, 
			created_at, 
			updated_at, 
			image, 
			item_id
		FROM item_image
		WHERE item_id = $1
		`
	rows, err := db.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item model.ItemImage
		err := rows.Scan(
			&item.Id,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.Image,
			&item.ItemId,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}
