package handlers

import (
	"StudentAPI/internal/models"
	"encoding/json"
	"net/http"
)

var grades []models.Grade

func CreateGrade(w http.ResponseWriter, r *http.Request) {
	var grade models.Grade
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&grade); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	grades = append(grades, grade)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(grade)
}
