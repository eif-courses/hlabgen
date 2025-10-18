package models

import "time"

type Booking struct {
	ID     string    `json:"id"`
	TripID string    `json:"trip_id"`
	UserID string    `json:"user_id"`
	Date   time.Time `json:"date"`
	Amount float64   `json:"amount"`
}
