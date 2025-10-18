package routes

import (
	"EventAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	r.HandleFunc("/events", handlers.CreateEvent).Methods("POST")
	r.HandleFunc("/events/{id:[0-9]+}", handlers.GetEvent).Methods("GET")
	r.HandleFunc("/events/{id:[0-9]+}", handlers.UpdateEvent).Methods("PUT")
	r.HandleFunc("/events/{id:[0-9]+}", handlers.DeleteEvent).Methods("DELETE")
	r.HandleFunc("/tickets", handlers.CreateTicket).Methods("POST")
	r.HandleFunc("/tickets/{id:[0-9]+}", handlers.GetTicket).Methods("GET")
	r.HandleFunc("/tickets/{id:[0-9]+}", handlers.UpdateTicket).Methods("PUT")
	r.HandleFunc("/tickets/{id:[0-9]+}", handlers.DeleteTicket).Methods("DELETE")
	r.HandleFunc("/attendees", handlers.CreateAttendee).Methods("POST")
	r.HandleFunc("/venues", handlers.CreateVenue).Methods("POST")
	r.HandleFunc("/organizers", handlers.CreateOrganizer).Methods("POST")
}
