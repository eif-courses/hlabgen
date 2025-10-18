package routes

import (
	"HealthAPI/internal/handlers"
	"github.com/gorilla/mux"
)

// Register registers all routes for the application.
func Register() {
	r.HandleFunc("/patients", handlers.CreatePatient).Methods("POST")
	r.HandleFunc("/patients", handlers.GetPatients).Methods("GET")
	r.HandleFunc("/doctors", handlers.CreateDoctor).Methods("POST")
	r.HandleFunc("/doctors", handlers.GetDoctors).Methods("GET")
	r.HandleFunc("/records", handlers.CreateRecord).Methods("POST")
	r.HandleFunc("/records", handlers.GetRecords).Methods("GET")
}
