package model

type Bid struct {
	Id        int64  `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	BidPrice  int64  `json:"bid_price"`
	ItemId    int64  `json:"item_id"`
	UserId    int64  `json:"user_id"`
}

type BidDetail struct {
	TotalBid   int64 `json:"total_bid"`
	HighestBid int64 `json:"highest_bid"`
	HistoryBid []Bid `json:"history_bid"`
}

type BidPost struct {
	BidPrice int64 `form:"bid_price"`
	ItemId   int64 `form:"item_id"`
	UserId   int64 `form:"user_id"`
}
