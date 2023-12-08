package model

type Item struct {
	Id           int64  `json:"id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	IsAuction    bool   `json:"is_auction"`
	IsActive     bool   `json:"is_active"`
	StartAuction string `json:"start_auct"`
	EndAuction   string `json:"end_auct"`
	StartPrice   int64  `json:"start_price"`
	SellerId     int64  `json:"-"`
}

type ItemImage struct {
	Id        int64  `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Image     string `json:"image"`
	ItemId    int64  `json:"item_id"`
}

type ItemDetail struct {
	Item
	Images []ItemImage `json:"images"`
	Seller Seller      `json:"seller"`
	Bid    BidDetail   `json:"bid"`
}

type ItemDetailPagination struct {
	Page       int64        `json:"page"`
	PerPage    int64        `json:"per_page"`
	PageCount  int64        `json:"page_count"`
	TotalCount int64        `json:"total_count"`
	Records    []ItemDetail `json:"records"`
}

type ItemPost struct {
	Name         string `form:"name"`
	Description  string `form:"description"`
	StartAuction string `form:"start_auct"`
	EndAuction   string `form:"end_auct"`
	StartPrice   int64  `form:"start_price"`
	SellerId     int64  `form:"seller_id"`
}
