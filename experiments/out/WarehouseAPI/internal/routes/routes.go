package routes

import (
	"WarehouseAPI/internal/handlers"
	"github.com/gorilla/mux"
)

// Register sets up the routes for the API.
func Register() {
	router.HandleFunc("/items", handlers.CreateItem).Methods("POST")
	router.HandleFunc("/items", handlers.GetItems).Methods("GET")
	router.HandleFunc("/locations", handlers.CreateLocation).Methods("POST")
	router.HandleFunc("/locations", handlers.GetLocations).Methods("GET")
	router.HandleFunc("/shipments", handlers.CreateShipment).Methods("POST")
	router.HandleFunc("/shipments", handlers.GetShipments).Methods("GET")
}
