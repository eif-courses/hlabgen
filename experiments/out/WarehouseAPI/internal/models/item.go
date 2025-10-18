package models

import "encoding/json"

// Item represents a product in the warehouse.
type Item struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Quantity   int    `json:"quantity"`
	LocationID int    `json:"location_id"`
}

// ToJSON converts an Item to JSON.
func (i *Item) ToJSON() ([]byte, error) {
	return json.Marshal(i)
}
