package models

type Invoice struct {
	ID     int     `json:"id"`
	Amount float64 `json:"amount"`
	Paid   bool    `json:"paid"`
}
