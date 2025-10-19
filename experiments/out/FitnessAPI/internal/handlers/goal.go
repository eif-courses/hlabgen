package handlers

import (
	"FitnessAPI/internal/models"
	"encoding/json"
	"net/http"
)

var goals []models.Goal

func CreateGoal(w http.ResponseWriter, r *http.Request) {
	var goal models.Goal
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&goal); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	goals = append(goals, goal)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(goal)
}
