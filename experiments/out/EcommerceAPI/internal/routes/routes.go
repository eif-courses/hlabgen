package routes

import (
	"EcommerceAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	// Product routes
	r.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	r.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", handlers.GetProduct).Methods("GET")
	r.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE")
	// User routes
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
	// Cart routes
	r.HandleFunc("/carts", handlers.CreateCart).Methods("POST")
	r.HandleFunc("/carts", handlers.GetCarts).Methods("GET")
	r.HandleFunc("/carts/{id}", handlers.GetCart).Methods("GET")
	r.HandleFunc("/carts/{id}", handlers.UpdateCart).Methods("PUT")
	r.HandleFunc("/carts/{id}", handlers.DeleteCart).Methods("DELETE")
	// Order routes
	r.HandleFunc("/orders", handlers.CreateOrder).Methods("POST")
	r.HandleFunc("/orders", handlers.GetOrders).Methods("GET")
	r.HandleFunc("/orders/{id}", handlers.GetOrder).Methods("GET")
	r.HandleFunc("/orders/{id}", handlers.UpdateOrder).Methods("PUT")
	r.HandleFunc("/orders/{id}", handlers.DeleteOrder).Methods("DELETE")
}
