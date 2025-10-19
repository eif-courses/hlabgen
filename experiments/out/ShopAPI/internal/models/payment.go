package models

type Payment struct {
	ID          int     `json:"id"`
	OrderID     int     `json:"order_id"`
	Amount      float64 `json:"amount"`
	PaymentDate string  `json:"payment_date"`
}
