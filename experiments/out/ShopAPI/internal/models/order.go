package models

type Order struct {
	ID         int       `json:"id"`
	CustomerID int       `json:"customer_id"`
	Products   []Product `json:"products"`
	Total      float64   `json:"total"`
	Status     string    `json:"status"`
}
