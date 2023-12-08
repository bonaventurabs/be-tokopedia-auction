package postgre

import (
	"context"

	"github.com/bonaventurabs/be-tokopedia-auction/model"
)

func (db *DB) GetBidByItemId(ctx context.Context, ItemId int64) ([]model.Bid, error) {
	var (
		items []model.Bid
	)

	query := `SELECT 
			id, 
			created_at, 
			updated_at, 
			bid_price,
			item_id,
			user_id
		FROM item_bid
		WHERE item_id = $1
		`
	rows, err := db.DB.QueryContext(ctx, query, ItemId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item model.Bid
		err := rows.Scan(
			&item.Id,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.BidPrice,
			&item.ItemId,
			&item.UserId,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}
