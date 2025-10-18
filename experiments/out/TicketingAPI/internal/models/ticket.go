package models

import "time"

type Ticket struct {
	ID         int       `json:"id"`
	EventID    int       `json:"event_id"`
	AttendeeID int       `json:"attendee_id"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
