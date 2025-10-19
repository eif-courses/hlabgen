package handlers

import (
	"ProjectAPI/internal/models"
	"encoding/json"
	"net/http"
)

var tasks []models.Task

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tasks = append(tasks, task)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
func GetTask(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single task
	w.WriteHeader(http.StatusOK)
}
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a task
	w.WriteHeader(http.StatusOK)
}
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a task
	w.WriteHeader(http.StatusNoContent)
}
