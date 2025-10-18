package models

import "encoding/json"

// Location represents a storage location in the warehouse.
type Location struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ToJSON converts a Location to JSON.
func (l *Location) ToJSON() ([]byte, error) {
	return json.Marshal(l)
}
