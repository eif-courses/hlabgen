package routes

import (
	"StudentAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register() {
	r.HandleFunc("/students", handlers.CreateStudent).Methods("POST")
	r.HandleFunc("/students", handlers.GetStudents).Methods("GET")
	r.HandleFunc("/students/{id}", handlers.UpdateStudent).Methods("PUT")
	r.HandleFunc("/students/{id}", handlers.DeleteStudent).Methods("DELETE")
	r.HandleFunc("/courses", handlers.CreateCourse).Methods("POST")
	r.HandleFunc("/courses", handlers.GetCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", handlers.UpdateCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", handlers.DeleteCourse).Methods("DELETE")
	r.HandleFunc("/grades", handlers.CreateGrade).Methods("POST")
	r.HandleFunc("/grades", handlers.GetGrades).Methods("GET")
	r.HandleFunc("/grades/{id}", handlers.UpdateGrade).Methods("PUT")
	r.HandleFunc("/grades/{id}", handlers.DeleteGrade).Methods("DELETE")
}
