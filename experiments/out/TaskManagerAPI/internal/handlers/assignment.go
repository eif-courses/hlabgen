package handlers

import (
	"TaskManagerAPI/internal/models"
	"encoding/json"
	"net/http"
)

var assignments []models.Assignment

func CreateAssignment(w http.ResponseWriter, r *http.Request) {
	var assignment models.Assignment
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&assignment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	assignments = append(assignments, assignment)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(assignment)
}
