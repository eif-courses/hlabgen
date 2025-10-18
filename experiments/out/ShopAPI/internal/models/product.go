package models

import "encoding/json"

// Product represents a product in the shop.
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// ToJSON converts a Product to JSON.
func (p *Product) ToJSON() ([]byte, error) {
	return json.Marshal(p)
}
