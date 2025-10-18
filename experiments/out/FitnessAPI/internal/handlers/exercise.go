package handlers

import (
	"FitnessAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateExercise() {
	var exercise models.Exercise
	if err := json.NewDecoder(r.Body).Decode(&exercise); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(exercise)
}

// Other CRUD functions for Exercise
