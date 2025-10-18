package models

import "time"

// Movie represents a movie entity.
type Movie struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	DirectorID  int       `json:"director_id"`
	ReleaseDate time.Time `json:"release_date"`
	Genre       string    `json:"genre"`
}
