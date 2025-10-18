package models

import "time"

type Invoice struct {
	ID        int       `json:"id"`
	AccountID int       `json:"account_id"`
	Amount    float64   `json:"amount"`
	DueDate   time.Time `json:"due_date"`
	CreatedAt time.Time `json:"created_at"`
}
