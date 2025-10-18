package routes

import (
	"InventoryAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register() {
	router.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	router.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	router.HandleFunc("/orders", handlers.CreateOrder).Methods("POST")
	router.HandleFunc("/orders", handlers.GetOrders).Methods("GET")
	router.HandleFunc("/suppliers", handlers.CreateSupplier).Methods("POST")
	router.HandleFunc("/suppliers", handlers.GetSuppliers).Methods("GET")
}
