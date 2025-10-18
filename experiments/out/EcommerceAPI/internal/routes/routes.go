package routes

import (
	"EcommerceAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register() {
	r.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	r.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	r.HandleFunc("/products/{id:[0-9]+}", handlers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id:[0-9]+}", handlers.DeleteProduct).Methods("DELETE")
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/carts", handlers.CreateCart).Methods("POST")
	r.HandleFunc("/orders", handlers.CreateOrder).Methods("POST")
}
