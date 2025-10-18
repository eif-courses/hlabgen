package models

import "time"

type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	VenueID     int       `json:"venue_id"`
	OrganizerID int       `json:"organizer_id"`
	Capacity    int       `json:"capacity"`
}
