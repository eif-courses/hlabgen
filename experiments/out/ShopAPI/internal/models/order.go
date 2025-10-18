package models

import "time"

// Order represents a customer's order.
type Order struct {
	ID         int       `json:"id"`
	CustomerID int       `json:"customer_id"`
	Total      float64   `json:"total"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
