package routes

import (
	"InventoryAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	// Product routes
	r.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	r.HandleFunc("/products", handlers.GetProducts).Methods("GET")

	// Order routes
	r.HandleFunc("/orders", handlers.CreateOrder).Methods("POST")
	r.HandleFunc("/orders", handlers.GetOrders).Methods("GET")

	// Supplier routes
	r.HandleFunc("/suppliers", handlers.CreateSupplier).Methods("POST")
	r.HandleFunc("/suppliers", handlers.GetSuppliers).Methods("GET")
}
