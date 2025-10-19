package routes

import (
	"HealthAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	// Patient routes
	r.HandleFunc("/patients", handlers.CreatePatient).Methods("POST")
	r.HandleFunc("/patients", handlers.GetPatients).Methods("GET")
	r.HandleFunc("/patients/{id}", handlers.GetPatient).Methods("GET")
	r.HandleFunc("/patients/{id}", handlers.UpdatePatient).Methods("PUT")
	r.HandleFunc("/patients/{id}", handlers.DeletePatient).Methods("DELETE")
}
