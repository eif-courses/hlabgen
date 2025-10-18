package models

import "encoding/json"

// Payment represents a payment in the shop.
type Payment struct {
	ID      int     `json:"id"`
	OrderID int     `json:"order_id"`
	Amount  float64 `json:"amount"`
	Status  string  `json:"status"`
}

// ToJSON converts a Payment to JSON.
func (p *Payment) ToJSON() ([]byte, error) {
	return json.Marshal(p)
}
