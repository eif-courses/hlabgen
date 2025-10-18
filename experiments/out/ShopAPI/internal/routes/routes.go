package routes

import (
	"ShopAPI/internal/handlers"
	"github.com/gorilla/mux"
)

// Register registers the routes for the application.
func Register() {
	r.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	r.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE")
	// Add routes for orders, customers, etc.
}
