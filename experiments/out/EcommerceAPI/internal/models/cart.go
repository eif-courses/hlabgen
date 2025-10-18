package models

import "encoding/json"

// Cart represents a shopping cart in the e-commerce system.
type Cart struct {
	ID       int       `json:"id"`
	UserID   int       `json:"user_id"`
	Products []Product `json:"products"`
}

// ToJSON converts a Cart to JSON.
func (c *Cart) ToJSON() ([]byte, error) {
	return json.Marshal(c)
}
