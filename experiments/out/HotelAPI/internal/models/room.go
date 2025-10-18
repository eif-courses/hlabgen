package models

import "time"

type Room struct {
	ID        int       `json:"id"`
	HotelID   int       `json:"hotel_id"`
	RoomType  string    `json:"room_type"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
