package routes

import (
	"HotelAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	r.HandleFunc("/hotels", handlers.CreateHotel).Methods("POST")
	r.HandleFunc("/hotels", handlers.GetHotels).Methods("GET")
	r.HandleFunc("/rooms", handlers.CreateRoom).Methods("POST")
	r.HandleFunc("/rooms", handlers.GetRooms).Methods("GET")
	r.HandleFunc("/bookings", handlers.CreateBooking).Methods("POST")
	r.HandleFunc("/bookings", handlers.GetBookings).Methods("GET")
	r.HandleFunc("/guests", handlers.CreateGuest).Methods("POST")
	r.HandleFunc("/guests", handlers.GetGuests).Methods("GET")
	r.HandleFunc("/reviews", handlers.CreateReview).Methods("POST")
	r.HandleFunc("/reviews", handlers.GetReviews).Methods("GET")
}
