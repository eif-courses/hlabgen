package models

type Cart struct {
	UserID int     `json:"user_id"`
	Items  []Item  `json:"items"`
	Total  float64 `json:"total"`
}
type Item struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
