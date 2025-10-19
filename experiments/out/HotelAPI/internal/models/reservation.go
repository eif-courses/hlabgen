package models

type Reservation struct {
	ID        int    `json:"id"`
	RoomID    int    `json:"room_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
