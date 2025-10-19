package models

type Cart struct {
	ID       int       `json:"id"`
	Customer Customer  `json:"customer"`
	Items    []Product `json:"items"`
}
