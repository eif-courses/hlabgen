package models

// Cart represents a shopping cart.
type Cart struct {
	ID         int       `json:"id"`
	CustomerID int       `json:"customer_id"`
	Items      []Product `json:"items"`
}
