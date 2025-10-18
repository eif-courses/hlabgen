package models

import "encoding/json"

// Restaurant represents a restaurant entity.
type Restaurant struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

// ToJSON converts a Restaurant to JSON.
func (r *Restaurant) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}
