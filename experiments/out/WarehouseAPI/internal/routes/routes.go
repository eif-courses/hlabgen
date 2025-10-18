package routes

import (
	"WarehouseAPI/internal/handlers"
	"github.com/gorilla/mux"
)

// Register registers the routes for the API.
func Register() {
	r.HandleFunc("/items", handlers.CreateItem).Methods("POST")
	r.HandleFunc("/items", handlers.GetItems).Methods("GET")
	r.HandleFunc("/items/{id}", handlers.UpdateItem).Methods("PUT")
	r.HandleFunc("/items/{id}", handlers.DeleteItem).Methods("DELETE")

	r.HandleFunc("/locations", handlers.CreateLocation).Methods("POST")
	r.HandleFunc("/locations", handlers.GetLocations).Methods("GET")
	r.HandleFunc("/locations/{id}", handlers.UpdateLocation).Methods("PUT")
	r.HandleFunc("/locations/{id}", handlers.DeleteLocation).Methods("DELETE")

	r.HandleFunc("/shipments", handlers.CreateShipment).Methods("POST")
	r.HandleFunc("/shipments", handlers.GetShipments).Methods("GET")
	r.HandleFunc("/shipments/{id}", handlers.UpdateShipment).Methods("PUT")
	r.HandleFunc("/shipments/{id}", handlers.DeleteShipment).Methods("DELETE")
}
