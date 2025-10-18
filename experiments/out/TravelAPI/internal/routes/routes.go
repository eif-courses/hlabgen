package routes

import (
	"TravelAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register() {
	r.HandleFunc("/trips", handlers.CreateTrip).Methods("POST")
	r.HandleFunc("/trips/{id}", handlers.GetTrip).Methods("GET")
	r.HandleFunc("/trips/{id}", handlers.UpdateTrip).Methods("PUT")
	r.HandleFunc("/trips/{id}", handlers.DeleteTrip).Methods("DELETE")
	r.HandleFunc("/bookings", handlers.CreateBooking).Methods("POST")
	r.HandleFunc("/bookings/{id}", handlers.GetBooking).Methods("GET")
	r.HandleFunc("/bookings/{id}", handlers.UpdateBooking).Methods("PUT")
	r.HandleFunc("/bookings/{id}", handlers.DeleteBooking).Methods("DELETE")
	r.HandleFunc("/destinations", handlers.CreateDestination).Methods("POST")
	r.HandleFunc("/destinations/{id}", handlers.GetDestination).Methods("GET")
	r.HandleFunc("/destinations/{id}", handlers.UpdateDestination).Methods("PUT")
	r.HandleFunc("/destinations/{id}", handlers.DeleteDestination).Methods("DELETE")
}
