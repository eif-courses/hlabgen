package routes

import (
	"SchoolAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	// Teacher routes
	r.HandleFunc("/teachers", handlers.CreateTeacher).Methods("POST")
	r.HandleFunc("/teachers", handlers.GetTeachers).Methods("GET")
	r.HandleFunc("/teachers/{id}", handlers.GetTeacher).Methods("GET")
	r.HandleFunc("/teachers/{id}", handlers.UpdateTeacher).Methods("PUT")
	r.HandleFunc("/teachers/{id}", handlers.DeleteTeacher).Methods("DELETE")
}
