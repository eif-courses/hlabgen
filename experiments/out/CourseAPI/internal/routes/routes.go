package routes

import (
	"CourseAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	// Course routes
	r.HandleFunc("/courses", handlers.CreateCourse).Methods("POST")
	r.HandleFunc("/courses", handlers.GetCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", handlers.GetCourse).Methods("GET")
	r.HandleFunc("/courses/{id}", handlers.UpdateCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", handlers.DeleteCourse).Methods("DELETE")
}
