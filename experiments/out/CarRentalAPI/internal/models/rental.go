package models

import "encoding/json"

// Rental represents a rental transaction in the system.
type Rental struct {
	ID         int     `json:"id"`
	CarID      int     `json:"car_id"`
	CustomerID int     `json:"customer_id"`
	StartDate  string  `json:"start_date"`
	EndDate    string  `json:"end_date"`
	TotalCost  float64 `json:"total_cost"`
}

// ToJSON converts a Rental to JSON.
func (r *Rental) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

// FromJSON converts JSON to a Rental.
