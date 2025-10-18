package models

import "encoding/json"

// Artist represents a music artist.
type Artist struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ToJSON converts an Artist to JSON.
func (a *Artist) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}
