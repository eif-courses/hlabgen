package models

import "time"

type Transaction struct {
	ID        int       `json:"id"`
	AccountID int       `json:"account_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
