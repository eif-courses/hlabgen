package handlers

import (
	"RecipeAPI/internal/models"
	"encoding/json"
	"net/http"
)

var steps []models.Step

func CreateStep(w http.ResponseWriter, r *http.Request) {
	var step models.Step
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&step); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	steps = append(steps, step)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(step)
}

func GetSteps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(steps)
}
func GetStep(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single step
	w.WriteHeader(http.StatusOK)
}
func UpdateStep(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a step
	w.WriteHeader(http.StatusOK)
}
func DeleteStep(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a step
	w.WriteHeader(http.StatusNoContent)
}
