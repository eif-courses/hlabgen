package routes

import (
	"CarRentalAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	r.HandleFunc("/cars", handlers.CreateCar).Methods("POST")
	r.HandleFunc("/customers", handlers.CreateCustomer).Methods("POST")
	r.HandleFunc("/rentals", handlers.CreateRental).Methods("POST")
}
