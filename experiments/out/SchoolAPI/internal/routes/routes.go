package routes

import (
	"SchoolAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register() {
	router.HandleFunc("/teachers", handlers.CreateTeacher).Methods("POST")
	router.HandleFunc("/teachers", handlers.GetTeachers).Methods("GET")
	router.HandleFunc("/teachers/{id}", handlers.UpdateTeacher).Methods("PUT")
	router.HandleFunc("/teachers/{id}", handlers.DeleteTeacher).Methods("DELETE")

	router.HandleFunc("/students", handlers.CreateStudent).Methods("POST")
	router.HandleFunc("/students", handlers.GetStudents).Methods("GET")
	router.HandleFunc("/students/{id}", handlers.UpdateStudent).Methods("PUT")
	router.HandleFunc("/students/{id}", handlers.DeleteStudent).Methods("DELETE")

	router.HandleFunc("/classes", handlers.CreateClass).Methods("POST")
	router.HandleFunc("/classes", handlers.GetClasses).Methods("GET")
	router.HandleFunc("/classes/{id}", handlers.UpdateClass).Methods("PUT")
	router.HandleFunc("/classes/{id}", handlers.DeleteClass).Methods("DELETE")
}
