package handlers

import (
	"FitnessAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateWorkout() {
	var workout models.Workout
	if err := json.NewDecoder(r.Body).Decode(&workout); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(workout)
}

// Other CRUD functions for Workout
