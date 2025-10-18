package models

import "encoding/json"

// Album represents a music album.
type Album struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	ArtistID int    `json:"artist_id"`
}

// ToJSON converts an Album to JSON.
func (a *Album) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}
