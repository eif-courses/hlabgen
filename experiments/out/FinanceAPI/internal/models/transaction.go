package models

type Transaction struct {
	ID          int     `json:"id"`
	AccountID   int     `json:"account_id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"created_at"`
}
