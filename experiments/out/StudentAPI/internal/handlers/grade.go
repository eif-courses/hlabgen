package handlers

import (
	"StudentAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateGrade() {
	var grade models.Grade
	if err := json.NewDecoder(r.Body).Decode(&grade); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(grade)
}

func GetGrades() {
	// Implementation for fetching grades
}
func UpdateGrade() {
	// Implementation for updating a grade
}
func DeleteGrade() {
	// Implementation for deleting a grade
}
