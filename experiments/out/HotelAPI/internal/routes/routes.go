package routes

import (
	"HotelAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	r.HandleFunc("/hotels", handlers.CreateHotel).Methods("POST")
	r.HandleFunc("/hotels", handlers.GetHotels).Methods("GET")
	r.HandleFunc("/hotels/{id}", handlers.UpdateHotel).Methods("PUT")
	r.HandleFunc("/hotels/{id}", handlers.DeleteHotel).Methods("DELETE")

	r.HandleFunc("/rooms", handlers.CreateRoom).Methods("POST")
	r.HandleFunc("/rooms", handlers.GetRooms).Methods("GET")
	r.HandleFunc("/rooms/{id}", handlers.UpdateRoom).Methods("PUT")
	r.HandleFunc("/rooms/{id}", handlers.DeleteRoom).Methods("DELETE")

	r.HandleFunc("/bookings", handlers.CreateBooking).Methods("POST")
	r.HandleFunc("/bookings", handlers.GetBookings).Methods("GET")
	r.HandleFunc("/bookings/{id}", handlers.UpdateBooking).Methods("PUT")
	r.HandleFunc("/bookings/{id}", handlers.DeleteBooking).Methods("DELETE")

	r.HandleFunc("/guests", handlers.CreateGuest).Methods("POST")
	r.HandleFunc("/guests", handlers.GetGuests).Methods("GET")
	r.HandleFunc("/guests/{id}", handlers.UpdateGuest).Methods("PUT")
	r.HandleFunc("/guests/{id}", handlers.DeleteGuest).Methods("DELETE")

	r.HandleFunc("/reviews", handlers.CreateReview).Methods("POST")
	r.HandleFunc("/reviews", handlers.GetReviews).Methods("GET")
	r.HandleFunc("/reviews/{id}", handlers.UpdateReview).Methods("PUT")
	r.HandleFunc("/reviews/{id}", handlers.DeleteReview).Methods("DELETE")
}
