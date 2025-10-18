package routes

import (
	"InventoryAPI/internal/handlers"
	"github.com/gorilla/mux"
)

// Register registers the routes for the application.
func Register() {
	r.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	r.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE")
	r.HandleFunc("/orders", handlers.CreateOrder).Methods("POST")
	r.HandleFunc("/orders", handlers.GetOrders).Methods("GET")
	r.HandleFunc("/orders/{id}", handlers.UpdateOrder).Methods("PUT")
	r.HandleFunc("/orders/{id}", handlers.DeleteOrder).Methods("DELETE")
	r.HandleFunc("/suppliers", handlers.CreateSupplier).Methods("POST")
	r.HandleFunc("/suppliers", handlers.GetSuppliers).Methods("GET")
	r.HandleFunc("/suppliers/{id}", handlers.UpdateSupplier).Methods("PUT")
	r.HandleFunc("/suppliers/{id}", handlers.DeleteSupplier).Methods("DELETE")
}
