package handlers

import (
	"FitnessAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateGoal() {
	var goal models.Goal
	if err := json.NewDecoder(r.Body).Decode(&goal); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(goal)
}

func GetGoal() {
	// Implementation for getting a goal
}
func UpdateGoal() {
	// Implementation for updating a goal
}
func DeleteGoal() {
	// Implementation for deleting a goal
}
