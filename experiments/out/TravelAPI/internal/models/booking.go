package models

type Booking struct {
	ID     int    `json:"id"`
	TripID int    `json:"trip_id"`
	Date   string `json:"date"`
	Status string `json:"status"`
}
