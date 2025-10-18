package routes

import (
	"github.com/gorilla/mux"
)

func Register() {
	r.HandleFunc("/tickets", handlers.CreateTicket).Methods("POST")
	r.HandleFunc("/tickets", handlers.GetTickets).Methods("GET")
	r.HandleFunc("/tickets/{id}", handlers.UpdateTicket).Methods("PUT")
	r.HandleFunc("/tickets/{id}", handlers.DeleteTicket).Methods("DELETE")
	// Add routes for events, attendees, and venues similarly
}
