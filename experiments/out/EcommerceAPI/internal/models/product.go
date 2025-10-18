package models

import "encoding/json"

// Product represents a product in the e-commerce system.
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

// ToJSON converts a Product to JSON.
func (p *Product) ToJSON() ([]byte, error) {
	return json.Marshal(p)
}
