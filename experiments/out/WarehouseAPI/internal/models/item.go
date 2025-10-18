package models

import "encoding/json"

// Item represents an item in the warehouse.
type Item struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Quantity   int    `json:"quantity"`
	LocationID int    `json:"location_id"`
}

// ToJSON converts the Item to JSON format.
func (i *Item) ToJSON() ([]byte, error) {
	return json.Marshal(i)
}
