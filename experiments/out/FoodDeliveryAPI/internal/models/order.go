package models

type Order struct {
	ID       int       `json:"id"`
	UserID   int       `json:"user_id"`
	Products []Product `json:"products"`
	Total    float64   `json:"total"`
}
