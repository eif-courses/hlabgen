package models

import "time"

type Song struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	ArtistID  int       `json:"artist_id"`
	AlbumID   int       `json:"album_id"`
	Duration  int       `json:"duration"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
