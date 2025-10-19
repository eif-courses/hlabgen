package routes

import (
	"TicketingAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	// Ticket routes
	r.HandleFunc("/tickets", handlers.CreateTicket).Methods("POST")
	r.HandleFunc("/tickets", handlers.GetTickets).Methods("GET")
	r.HandleFunc("/tickets/{id}", handlers.GetTicket).Methods("GET")
	r.HandleFunc("/tickets/{id}", handlers.UpdateTicket).Methods("PUT")
	r.HandleFunc("/tickets/{id}", handlers.DeleteTicket).Methods("DELETE")

	// Event routes
	r.HandleFunc("/events", handlers.CreateEvent).Methods("POST")
	r.HandleFunc("/events", handlers.GetEvents).Methods("GET")
	r.HandleFunc("/events/{id}", handlers.GetEvent).Methods("GET")
	r.HandleFunc("/events/{id}", handlers.UpdateEvent).Methods("PUT")
	r.HandleFunc("/events/{id}", handlers.DeleteEvent).Methods("DELETE")

	// Attendee routes
	r.HandleFunc("/attendees", handlers.CreateAttendee).Methods("POST")
	r.HandleFunc("/attendees", handlers.GetAttendees).Methods("GET")
	r.HandleFunc("/attendees/{id}", handlers.GetAttendee).Methods("GET")
	r.HandleFunc("/attendees/{id}", handlers.UpdateAttendee).Methods("PUT")
	r.HandleFunc("/attendees/{id}", handlers.DeleteAttendee).Methods("DELETE")

	// Venue routes
	r.HandleFunc("/venues", handlers.CreateVenue).Methods("POST")
	r.HandleFunc("/venues", handlers.GetVenues).Methods("GET")
	r.HandleFunc("/venues/{id}", handlers.GetVenue).Methods("GET")
	r.HandleFunc("/venues/{id}", handlers.UpdateVenue).Methods("PUT")
	r.HandleFunc("/venues/{id}", handlers.DeleteVenue).Methods("DELETE")
}
