package models

import "time"

type Rental struct {
	ID         int       `json:"id"`
	CarID      int       `json:"car_id"`
	CustomerID int       `json:"customer_id"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	TotalCost  float64   `json:"total_cost"`
}
