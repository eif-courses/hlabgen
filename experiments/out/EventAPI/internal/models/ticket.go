package models

type Ticket struct {
	ID      int     `json:"id"`
	EventID int     `json:"event_id"`
	Price   float64 `json:"price"`
	Status  string  `json:"status"`
}
