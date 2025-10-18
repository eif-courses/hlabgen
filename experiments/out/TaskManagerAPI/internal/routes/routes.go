package routes

import (
	"TaskManagerAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register() {
	r.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")
	r.HandleFunc("/teams", handlers.CreateTeam).Methods("POST")
	r.HandleFunc("/assignments", handlers.CreateAssignment).Methods("POST")
}
