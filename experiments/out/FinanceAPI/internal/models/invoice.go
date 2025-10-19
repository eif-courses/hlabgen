package models

type Invoice struct {
	ID     int     `json:"id"`
	Client string  `json:"client"`
	Amount float64 `json:"amount"`
	Paid   bool    `json:"paid"`
}
