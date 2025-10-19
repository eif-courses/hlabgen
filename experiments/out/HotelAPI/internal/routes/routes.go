package routes

import (
	"HotelAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	// Hotel routes
	r.HandleFunc("/hotels", handlers.CreateHotel).Methods("POST")
	r.HandleFunc("/rooms", handlers.CreateRoom).Methods("POST")
	r.HandleFunc("/reservations", handlers.CreateReservation).Methods("POST")
}
