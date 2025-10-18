package models

type Cart struct {
	ID         int       `json:"id"`
	CustomerID int       `json:"customer_id"`
	Products   []Product `json:"products"`
}
