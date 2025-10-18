package routes

import (
	"FoodDeliveryAPI/internal/handlers"
	"github.com/gorilla/mux"
)

// Register registers the routes for the application.
func Register() {
	r.HandleFunc("/restaurants", handlers.CreateRestaurant).Methods("POST")
	r.HandleFunc("/menuitems", handlers.CreateMenuItem).Methods("POST")
	r.HandleFunc("/orders", handlers.CreateOrder).Methods("POST")
}
