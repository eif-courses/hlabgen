package models

import "time"

type Reservation struct {
	ID        int       `json:"id"`
	RoomID    int       `json:"room_id"`
	GuestName string    `json:"guest_name"`
	CheckIn   time.Time `json:"check_in"`
	CheckOut  time.Time `json:"check_out"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
