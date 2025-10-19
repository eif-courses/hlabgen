package routes

import (
	"CarRentalAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	// Car routes
	r.HandleFunc("/cars", handlers.CreateCar).Methods("POST")
	r.HandleFunc("/cars", handlers.GetCars).Methods("GET")
	r.HandleFunc("/cars/{id}", handlers.GetCar).Methods("GET")
	r.HandleFunc("/cars/{id}", handlers.UpdateCar).Methods("PUT")
	r.HandleFunc("/cars/{id}", handlers.DeleteCar).Methods("DELETE")

	// Customer routes
	r.HandleFunc("/customers", handlers.CreateCustomer).Methods("POST")
	r.HandleFunc("/customers", handlers.GetCustomers).Methods("GET")
	r.HandleFunc("/customers/{id}", handlers.GetCustomer).Methods("GET")
	r.HandleFunc("/customers/{id}", handlers.UpdateCustomer).Methods("PUT")
	r.HandleFunc("/customers/{id}", handlers.DeleteCustomer).Methods("DELETE")

	// Rental routes
	r.HandleFunc("/rentals", handlers.CreateRental).Methods("POST")
	r.HandleFunc("/rentals", handlers.GetRentals).Methods("GET")
	r.HandleFunc("/rentals/{id}", handlers.GetRental).Methods("GET")
	r.HandleFunc("/rentals/{id}", handlers.UpdateRental).Methods("PUT")
	r.HandleFunc("/rentals/{id}", handlers.DeleteRental).Methods("DELETE")
}
