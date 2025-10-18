package routes

import (
	"ClinicAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register() {
	r.HandleFunc("/doctors", handlers.CreateDoctor).Methods("POST")
	r.HandleFunc("/doctors", handlers.GetDoctors).Methods("GET")
	r.HandleFunc("/doctors/{id}", handlers.UpdateDoctor).Methods("PUT")
	r.HandleFunc("/doctors/{id}", handlers.DeleteDoctor).Methods("DELETE")

	r.HandleFunc("/patients", handlers.CreatePatient).Methods("POST")
	r.HandleFunc("/patients", handlers.GetPatients).Methods("GET")
	r.HandleFunc("/patients/{id}", handlers.UpdatePatient).Methods("PUT")
	r.HandleFunc("/patients/{id}", handlers.DeletePatient).Methods("DELETE")

	r.HandleFunc("/appointments", handlers.CreateAppointment).Methods("POST")
	r.HandleFunc("/appointments", handlers.GetAppointments).Methods("GET")
	r.HandleFunc("/appointments/{id}", handlers.UpdateAppointment).Methods("PUT")
	r.HandleFunc("/appointments/{id}", handlers.DeleteAppointment).Methods("DELETE")
}
