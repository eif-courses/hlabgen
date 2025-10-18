package models

import "time"

type Hotel struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Location string  `json:"location"`
	Rating   float64 `json:"rating"`
}

type Room struct {
	ID        int     `json:"id"`
	HotelID   int     `json:"hotel_id"`
	RoomType  string  `json:"room_type"`
	Price     float64 `json:"price"`
	Available bool    `json:"available"`
}

type Booking struct {
	ID        int       `json:"id"`
	RoomID    int       `json:"room_id"`
	GuestID   int       `json:"guest_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type Guest struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Review struct {
	ID      int     `json:"id"`
	HotelID int     `json:"hotel_id"`
	GuestID int     `json:"guest_id"`
	Rating  float64 `json:"rating"`
	Comment string  `json:"comment"`
}
