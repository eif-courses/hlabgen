package models

import "encoding/json"

// Order represents an order in the e-commerce system.
type Order struct {
	ID       int       `json:"id"`
	UserID   int       `json:"user_id"`
	Products []Product `json:"products"`
	Total    float64   `json:"total"`
}

// ToJSON converts an Order to JSON.
func (o *Order) ToJSON() ([]byte, error) {
	return json.Marshal(o)
}
