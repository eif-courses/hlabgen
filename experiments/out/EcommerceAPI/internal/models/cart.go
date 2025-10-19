package models

type Cart struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	Items  []struct {
		ProductID int `json:"product_id"`
		Quantity  int `json:"quantity"`
	} `json:"items"`
}
