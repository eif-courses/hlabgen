package models

type Order struct {
	UserID   int       `json:"userID"`
	Products []Product `json:"products"`
	Total    float64   `json:"total"`
}
