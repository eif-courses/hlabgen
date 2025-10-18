package handlers

import (
	"TaskManagerAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateAssignment() {
	var assignment models.Assignment
	if err := json.NewDecoder(r.Body).Decode(&assignment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Logic to save assignment to database
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(assignment)
}

func GetAssignments() {
	// Logic to retrieve assignments from database
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode([]models.Assignment{})
}
