package routes

import (
	"CourseAPI/internal/handlers"
	"github.com/gorilla/mux"
)

// Register registers the routes for the application.
func Register() {
	r.HandleFunc("/courses", handlers.CreateCourse).Methods("POST")
	r.HandleFunc("/courses", handlers.GetCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", handlers.UpdateCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", handlers.DeleteCourse).Methods("DELETE")
	r.HandleFunc("/enrollments", handlers.CreateEnrollment).Methods("POST")
	r.HandleFunc("/enrollments", handlers.GetEnrollments).Methods("GET")
}
