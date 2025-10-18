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

func GetExercise() {
	// Implementation for getting an exercise
}
func UpdateExercise() {
	// Implementation for updating an exercise
}
func DeleteExercise() {
	// Implementation for deleting an exercise
}
