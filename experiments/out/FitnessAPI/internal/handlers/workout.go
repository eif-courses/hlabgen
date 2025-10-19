package handlers

import (
	"FitnessAPI/internal/models"
	"encoding/json"
	"net/http"
)

var workouts []models.Workout

func CreateWorkout(w http.ResponseWriter, r *http.Request) {
	var workout models.Workout
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&workout); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	workouts = append(workouts, workout)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(workout)
}
