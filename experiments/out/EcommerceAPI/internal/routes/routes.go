package routes

import (
	"EcommerceAPI/internal/handlers"
	"github.com/gorilla/mux"
)

// Register registers all routes for the application.
func Register() {
	r.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	r.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	// Add more routes as needed
}
