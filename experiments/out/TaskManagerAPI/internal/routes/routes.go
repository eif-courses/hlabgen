package routes

import (
	"TaskManagerAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	// Task routes
	r.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", handlers.GetTask).Methods("GET")
	r.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	// Team routes
	r.HandleFunc("/teams", handlers.CreateTeam).Methods("POST")
	r.HandleFunc("/teams", handlers.GetTeams).Methods("GET")
	r.HandleFunc("/teams/{id}", handlers.GetTeam).Methods("GET")
	r.HandleFunc("/teams/{id}", handlers.UpdateTeam).Methods("PUT")
	r.HandleFunc("/teams/{id}", handlers.DeleteTeam).Methods("DELETE")

	// Assignment routes
	r.HandleFunc("/assignments", handlers.CreateAssignment).Methods("POST")
	r.HandleFunc("/assignments", handlers.GetAssignments).Methods("GET")
	r.HandleFunc("/assignments/{id}", handlers.GetAssignment).Methods("GET")
	r.HandleFunc("/assignments/{id}", handlers.UpdateAssignment).Methods("PUT")
	r.HandleFunc("/assignments/{id}", handlers.DeleteAssignment).Methods("DELETE")
}
