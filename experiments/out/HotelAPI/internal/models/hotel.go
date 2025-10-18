package models

import "time"

// Hotel represents a hotel entity.
type Hotel struct {
    ID          int       `json:"id"`
    Name        string    `json:"name"`
    Location    string    `json:"location"`
    Rating      float64   `json:"rating"`
    Description string    `json:"description"`
}

// Room represents a room entity.
type Room struct {
    ID          int       `json:"id"`
    HotelID     int       `json:"hotel_id"`
    RoomType    string    `json:"room_type"`
    Price       float64   `json:"price"`
    Availability bool      `json:"availability"`
}

// Booking represents a booking entity.
type Booking struct {
    ID        int       `json:"id"`
    RoomID    int       `json:"room_id"`
    GuestID   int       `json:"guest_id"`
    StartDate time.Time  `json:"start_date"`
    EndDate   time.Time  `json:"end_date"`
}

// Guest represents a guest entity.
type Guest struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Phone    string `json:"phone"`
}

// Review represents a review entity.
type Review struct {
    ID      int     `json:"id"`
    HotelID int     `json:"hotel_id"`
    GuestID int     `json:"guest_id"`
    Rating  float64 `json:"rating"`
    Comment string  `json:"comment"`
}