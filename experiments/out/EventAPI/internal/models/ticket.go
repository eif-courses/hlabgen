package models

import "time"

type Ticket struct {
	ID         int       `json:"id"`
	EventID    int       `json:"event_id"`
	AttendeeID int       `json:"attendee_id"`
	Price      float64   `json:"price"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}
