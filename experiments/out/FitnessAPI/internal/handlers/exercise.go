package handlers

import (
	"FitnessAPI/internal/models"
	"encoding/json"
	"net/http"
)

var exercises []models.Exercise

func CreateExercise(w http.ResponseWriter, r *http.Request) {
	var exercise models.Exercise
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&exercise); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	exercises = append(exercises, exercise)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(exercise)
}
