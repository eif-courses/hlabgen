package models

type Restaurant struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Rating  float64 `json:"rating"`
}
