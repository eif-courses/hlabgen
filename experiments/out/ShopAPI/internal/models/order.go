package models

type Order struct {
	ID       int       `json:"id"`
	Customer Customer  `json:"customer"`
	Products []Product `json:"products"`
	Total    float64   `json:"total"`
}
