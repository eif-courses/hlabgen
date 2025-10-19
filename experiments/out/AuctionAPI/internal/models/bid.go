package models

type Bid struct {
	ID        int     `json:"id"`
	AuctionID int     `json:"auction_id"`
	UserID    int     `json:"user_id"`
	Amount    float64 `json:"amount"`
}
