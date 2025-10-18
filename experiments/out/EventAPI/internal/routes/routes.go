package routes

import (
	"EventAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	r.HandleFunc("/events", handlers.CreateEvent).Methods("POST")
	r.HandleFunc("/events/{id}", handlers.GetEvent).Methods("GET")
	r.HandleFunc("/events/{id}", handlers.UpdateEvent).Methods("PUT")
	r.HandleFunc("/events/{id}", handlers.DeleteEvent).Methods("DELETE")

	r.HandleFunc("/tickets", handlers.CreateTicket).Methods("POST")
	r.HandleFunc("/tickets/{id}", handlers.GetTicket).Methods("GET")
	r.HandleFunc("/tickets/{id}", handlers.UpdateTicket).Methods("PUT")
	r.HandleFunc("/tickets/{id}", handlers.DeleteTicket).Methods("DELETE")

	r.HandleFunc("/attendees", handlers.CreateAttendee).Methods("POST")
	r.HandleFunc("/attendees/{id}", handlers.GetAttendee).Methods("GET")
	r.HandleFunc("/attendees/{id}", handlers.UpdateAttendee).Methods("PUT")
	r.HandleFunc("/attendees/{id}", handlers.DeleteAttendee).Methods("DELETE")

	r.HandleFunc("/venues", handlers.CreateVenue).Methods("POST")
	r.HandleFunc("/venues/{id}", handlers.GetVenue).Methods("GET")
	r.HandleFunc("/venues/{id}", handlers.UpdateVenue).Methods("PUT")
	r.HandleFunc("/venues/{id}", handlers.DeleteVenue).Methods("DELETE")

	r.HandleFunc("/organizers", handlers.CreateOrganizer).Methods("POST")
	r.HandleFunc("/organizers/{id}", handlers.GetOrganizer).Methods("GET")
	r.HandleFunc("/organizers/{id}", handlers.UpdateOrganizer).Methods("PUT")
	r.HandleFunc("/organizers/{id}", handlers.DeleteOrganizer).Methods("DELETE")
}
