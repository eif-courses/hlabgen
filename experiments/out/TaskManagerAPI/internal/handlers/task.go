package handlers

import (
	"TaskManagerAPI/internal/models"
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
