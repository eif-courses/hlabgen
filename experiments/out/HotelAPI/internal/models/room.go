package models

type Room struct {
	ID       int    `json:"id"`
	Number   string `json:"number"`
	Capacity int    `json:"capacity"`
	HotelID  int    `json:"hotel_id"`
}
