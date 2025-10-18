package models

import "time"

type Event struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	VenueID   int       `json:"venue_id"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
