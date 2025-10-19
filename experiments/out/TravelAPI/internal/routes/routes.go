package routes

import (
	"TravelAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	// Trip routes
	r.HandleFunc("/trips", handlers.CreateTrip).Methods("POST")
	r.HandleFunc("/trips", handlers.GetTrips).Methods("GET")
	// Booking routes
	r.HandleFunc("/bookings", handlers.CreateBooking).Methods("POST")
	r.HandleFunc("/bookings", handlers.GetBookings).Methods("GET")
	// Destination routes
	r.HandleFunc("/destinations", handlers.CreateDestination).Methods("POST")
	r.HandleFunc("/destinations", handlers.GetDestinations).Methods("GET")
}
