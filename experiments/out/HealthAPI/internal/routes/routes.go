package routes

import (
	"HealthAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register() {
	r.HandleFunc("/patients", handlers.CreatePatient).Methods("POST")
	r.HandleFunc("/patients", handlers.GetPatients).Methods("GET")
	r.HandleFunc("/patients/{id}", handlers.UpdatePatient).Methods("PUT")
	r.HandleFunc("/patients/{id}", handlers.DeletePatient).Methods("DELETE")
	r.HandleFunc("/doctors", handlers.CreateDoctor).Methods("POST")
	r.HandleFunc("/doctors", handlers.GetDoctors).Methods("GET")
	r.HandleFunc("/doctors/{id}", handlers.UpdateDoctor).Methods("PUT")
	r.HandleFunc("/doctors/{id}", handlers.DeleteDoctor).Methods("DELETE")
	r.HandleFunc("/records", handlers.CreateRecord).Methods("POST")
	r.HandleFunc("/records", handlers.GetRecords).Methods("GET")
	r.HandleFunc("/records/{id}", handlers.UpdateRecord).Methods("PUT")
	r.HandleFunc("/records/{id}", handlers.DeleteRecord).Methods("DELETE")
}
