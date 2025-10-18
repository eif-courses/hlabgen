package handlers

import (
	"TaskAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting tasks
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a task
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a task
}
