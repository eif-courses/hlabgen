package models

import "time"

type Trip struct {
	ID          string    `json:"id"`
	Destination string    `json:"destination"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Price       float64   `json:"price"`
}
