package models

type Invoice struct {
	ID        int     `json:"id"`
	AccountID int     `json:"account_id"`
	Amount    float64 `json:"amount"`
	DueDate   string  `json:"due_date"`
	Status    string  `json:"status"`
}
