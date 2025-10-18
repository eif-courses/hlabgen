package routes

import (
	"SchoolAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register() {
	r.HandleFunc("/teachers", handlers.CreateTeacher).Methods("POST")
	r.HandleFunc("/teachers", handlers.GetTeachers).Methods("GET")
	r.HandleFunc("/teachers/{id}", handlers.UpdateTeacher).Methods("PUT")
	r.HandleFunc("/teachers/{id}", handlers.DeleteTeacher).Methods("DELETE")

	r.HandleFunc("/students", handlers.CreateStudent).Methods("POST")
	r.HandleFunc("/students", handlers.GetStudents).Methods("GET")
	r.HandleFunc("/students/{id}", handlers.UpdateStudent).Methods("PUT")
	r.HandleFunc("/students/{id}", handlers.DeleteStudent).Methods("DELETE")

	r.HandleFunc("/classes", handlers.CreateClass).Methods("POST")
	r.HandleFunc("/classes", handlers.GetClasses).Methods("GET")
	r.HandleFunc("/classes/{id}", handlers.UpdateClass).Methods("PUT")
	r.HandleFunc("/classes/{id}", handlers.DeleteClass).Methods("DELETE")
}
