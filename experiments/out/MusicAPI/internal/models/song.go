package models

import "encoding/json"

// Song represents a music song.
type Song struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	AlbumID int    `json:"album_id"`
}

// ToJSON converts a Song to JSON.
func (s *Song) ToJSON() ([]byte, error) {
	return json.Marshal(s)
}
