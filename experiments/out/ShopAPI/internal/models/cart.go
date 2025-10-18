package models

// Cart represents a shopping cart for a customer.
type Cart struct {
	ID         int   `json:"id"`
	CustomerID int   `json:"customer_id"`
	Products   []int `json:"products"`
}
