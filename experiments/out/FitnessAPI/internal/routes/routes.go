package routes

import (
	"FitnessAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register() {
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")
	// Add routes for Workout, Exercise, Goal
}
