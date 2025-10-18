package handlers

import (
	"StudentAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateStudent() {
	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

func GetStudents() {
	// Implementation for fetching students
}
func UpdateStudent() {
	// Implementation for updating a student
}
func DeleteStudent() {
	// Implementation for deleting a student
}
