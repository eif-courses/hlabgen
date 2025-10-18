package models

import "encoding/json"

// Cart represents a shopping cart in the shop.
type Cart struct {
	ID         int       `json:"id"`
	CustomerID int       `json:"customer_id"`
	Products   []Product `json:"products"`
}

// ToJSON converts a Cart to JSON.
func (c *Cart) ToJSON() ([]byte, error) {
	return json.Marshal(c)
}
