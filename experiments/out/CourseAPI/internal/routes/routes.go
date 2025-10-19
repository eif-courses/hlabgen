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
	// Lesson routes
	r.HandleFunc("/lessons", handlers.CreateLesson).Methods("POST")
	r.HandleFunc("/lessons", handlers.GetLessons).Methods("GET")
	r.HandleFunc("/lessons/{id}", handlers.GetLesson).Methods("GET")
	r.HandleFunc("/lessons/{id}", handlers.UpdateLesson).Methods("PUT")
	r.HandleFunc("/lessons/{id}", handlers.DeleteLesson).Methods("DELETE")
	// Enrollment routes
	r.HandleFunc("/enrollments", handlers.CreateEnrollment).Methods("POST")
	r.HandleFunc("/enrollments", handlers.GetEnrollments).Methods("GET")
	r.HandleFunc("/enrollments/{id}", handlers.GetEnrollment).Methods("GET")
	r.HandleFunc("/enrollments/{id}", handlers.UpdateEnrollment).Methods("PUT")
	r.HandleFunc("/enrollments/{id}", handlers.DeleteEnrollment).Methods("DELETE")
}
