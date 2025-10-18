package routes

import (
	"TicketingAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register() {
	r.HandleFunc("/tickets", handlers.CreateTicket).Methods("POST")
	r.HandleFunc("/events", handlers.CreateEvent).Methods("POST")
	r.HandleFunc("/attendees", handlers.CreateAttendee).Methods("POST")
	r.HandleFunc("/venues", handlers.CreateVenue).Methods("POST")
}
