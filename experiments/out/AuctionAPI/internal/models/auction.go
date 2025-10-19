package models

type Auction struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
}
