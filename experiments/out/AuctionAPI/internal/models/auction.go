package models

type Auction struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	UserID      int     `json:"user_id"`
	StartPrice  float64 `json:"start_price"`
	EndTime     string  `json:"end_time"`
}
