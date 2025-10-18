package models

import "encoding/json"

// Customer represents a customer in the shop.
type Customer struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ToJSON converts a Customer to JSON.
func (c *Customer) ToJSON() ([]byte, error) {
	return json.Marshal(c)
}
