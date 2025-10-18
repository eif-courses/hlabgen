package models

import "time"

type Album struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	ArtistID    int       `json:"artist_id"`
	ReleaseDate time.Time `json:"release_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
