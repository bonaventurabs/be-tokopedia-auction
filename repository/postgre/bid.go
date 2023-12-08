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

func (db *DB) PostBid(ctx context.Context, bid model.BidPost) (model.Bid, error) {
	var (
		item model.Bid
	)

	query := `INSERT INTO item_bid (created_at, updated_at, bid_price, item_id, user_id) 
		VALUES (NOW(), NOW(), $1, $2, $3) RETURNING id, created_at, updated_at, bid_price, item_id, user_id`
	err := db.DB.QueryRowContext(ctx, query,
		bid.BidPrice,
		bid.ItemId,
		bid.UserId,
	).Scan(
		&item.Id,
		&item.CreatedAt,
		&item.UpdatedAt,
		&item.BidPrice,
		&item.ItemId,
		&item.UserId,
	)
	if err != nil {
		return model.Bid{}, err
	}

	return item, nil
}

func (db *DB) GetHighestBidByItemId(ctx context.Context, ItemId int64) (model.Bid, error) {
	var (
		item model.Bid
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
		ORDER BY bid_price DESC
		LIMIT 1
		`
	err := db.DB.QueryRowContext(ctx, query, ItemId).Scan(
		&item.Id,
		&item.CreatedAt,
		&item.UpdatedAt,
		&item.BidPrice,
		&item.ItemId,
		&item.UserId,
	)
	if err != nil {
		return model.Bid{}, err
	}

	return item, nil
}
