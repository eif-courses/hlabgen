package models

import "encoding/json"

// Order represents an order in the shop.
type Order struct {
	ID         int     `json:"id"`
	CustomerID int     `json:"customer_id"`
	Total      float64 `json:"total"`
}

// ToJSON converts an Order to JSON.
func (o *Order) ToJSON() ([]byte, error) {
	return json.Marshal(o)
}
