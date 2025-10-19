package routes

import (
	"ProjectAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	// Project routes
	r.HandleFunc("/projects", handlers.CreateProject).Methods("POST")
	r.HandleFunc("/projects", handlers.GetProjects).Methods("GET")
	r.HandleFunc("/projects/{id}", handlers.GetProject).Methods("GET")
	r.HandleFunc("/projects/{id}", handlers.UpdateProject).Methods("PUT")
	r.HandleFunc("/projects/{id}", handlers.DeleteProject).Methods("DELETE")

	// Task routes
	r.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", handlers.GetTask).Methods("GET")
	r.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	// Milestone routes
	r.HandleFunc("/milestones", handlers.CreateMilestone).Methods("POST")
	r.HandleFunc("/milestones", handlers.GetMilestones).Methods("GET")
	r.HandleFunc("/milestones/{id}", handlers.GetMilestone).Methods("GET")
	r.HandleFunc("/milestones/{id}", handlers.UpdateMilestone).Methods("PUT")
	r.HandleFunc("/milestones/{id}", handlers.DeleteMilestone).Methods("DELETE")
}
