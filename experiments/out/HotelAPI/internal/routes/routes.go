package routes

import (
	"HotelAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register() {
	r.HandleFunc("/hotels", handlers.CreateHotel).Methods("POST")
	r.HandleFunc("/hotels", handlers.GetHotels).Methods("GET")
	r.HandleFunc("/hotels/{id}", handlers.UpdateHotel).Methods("PUT")
	r.HandleFunc("/hotels/{id}", handlers.DeleteHotel).Methods("DELETE")
	r.HandleFunc("/rooms", handlers.CreateRoom).Methods("POST")
	r.HandleFunc("/rooms", handlers.GetRooms).Methods("GET")
	r.HandleFunc("/rooms/{id}", handlers.UpdateRoom).Methods("PUT")
	r.HandleFunc("/rooms/{id}", handlers.DeleteRoom).Methods("DELETE")
	r.HandleFunc("/reservations", handlers.CreateReservation).Methods("POST")
	r.HandleFunc("/reservations", handlers.GetReservations).Methods("GET")
	r.HandleFunc("/reservations/{id}", handlers.UpdateReservation).Methods("PUT")
	r.HandleFunc("/reservations/{id}", handlers.DeleteReservation).Methods("DELETE")
}
