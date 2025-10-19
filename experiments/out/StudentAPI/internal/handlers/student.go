package handlers

import (
	"StudentAPI/internal/models"
	"encoding/json"
	"net/http"
)

var students []models.Student

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	students = append(students, student)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}
