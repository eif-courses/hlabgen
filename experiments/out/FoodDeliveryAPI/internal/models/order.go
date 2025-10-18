package models

import "encoding/json"

// Order represents an order entity.
type Order struct {
	ID         int     `json:"id"`
	UserID     int     `json:"user_id"`
	MenuItemID int     `json:"menu_item_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}

// ToJSON converts an Order to JSON.
func (o *Order) ToJSON() ([]byte, error) {
	return json.Marshal(o)
}
