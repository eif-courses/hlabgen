package models

// Hotel represents a hotel entity.
type Hotel struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

// Room represents a room entity.
type Room struct {
	ID         int     `json:"id"`
	HotelID    int     `json:"hotel_id"`
	RoomNumber string  `json:"room_number"`
	Price      float64 `json:"price"`
}

// Reservation represents a reservation entity.
type Reservation struct {
	ID        int    `json:"id"`
	RoomID    int    `json:"room_id"`
	GuestName string `json:"guest_name"`
	CheckIn   string `json:"check_in"`
	CheckOut  string `json:"check_out"`
}
