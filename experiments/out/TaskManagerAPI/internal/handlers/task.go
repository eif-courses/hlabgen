package handlers

import (
	"TaskManagerAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateTask() {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Logic to save task to database
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func GetTasks() {
	// Logic to retrieve tasks from database
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode([]models.Task{})
}
func UpdateTask() {
	// Logic to update task in database
	w.WriteHeader(http.StatusOK)
}
func DeleteTask() {
	// Logic to delete task from database
	w.WriteHeader(http.StatusNoContent)
}
