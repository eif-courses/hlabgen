package routes

import (
	"FoodDeliveryAPI/internal/handlers"
	"github.com/gorilla/mux"
)

// Register registers all routes for the application.
func Register() {
	r.HandleFunc("/restaurants", handlers.CreateRestaurant).Methods("POST")
	r.HandleFunc("/restaurants", handlers.GetRestaurants).Methods("GET")
	r.HandleFunc("/restaurants/{id}", handlers.UpdateRestaurant).Methods("PUT")
	r.HandleFunc("/restaurants/{id}", handlers.DeleteRestaurant).Methods("DELETE")

	r.HandleFunc("/menuitems", handlers.CreateMenuItem).Methods("POST")
	r.HandleFunc("/menuitems", handlers.GetMenuItems).Methods("GET")
	r.HandleFunc("/menuitems/{id}", handlers.UpdateMenuItem).Methods("PUT")
	r.HandleFunc("/menuitems/{id}", handlers.DeleteMenuItem).Methods("DELETE")

	r.HandleFunc("/orders", handlers.CreateOrder).Methods("POST")
	r.HandleFunc("/orders", handlers.GetOrders).Methods("GET")
	r.HandleFunc("/orders/{id}", handlers.UpdateOrder).Methods("PUT")
	r.HandleFunc("/orders/{id}", handlers.DeleteOrder).Methods("DELETE")
}
