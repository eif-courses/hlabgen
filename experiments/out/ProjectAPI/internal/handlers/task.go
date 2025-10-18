package handlers

import (
	"ProjectAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateTask() {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func GetTasks() {
	// Implementation for retrieving tasks
}
func UpdateTask() {
	// Implementation for updating a task
}
func DeleteTask() {
	// Implementation for deleting a task
}
