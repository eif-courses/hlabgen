package models

type Destination struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Country     string  `json:"country"`
	Price       float64 `json:"price"`
}
