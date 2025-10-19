package models

type Hotel struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Rating  int    `json:"rating"`
}
