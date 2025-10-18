package handlers

import (
	"RecipeAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateStep() {
	var step models.Step
	if err := json.NewDecoder(r.Body).Decode(&step); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(step)
}

func GetSteps() {
	// Implementation for getting steps
}
func UpdateStep() {
	// Implementation for updating a step
}
func DeleteStep() {
	// Implementation for deleting a step
}
