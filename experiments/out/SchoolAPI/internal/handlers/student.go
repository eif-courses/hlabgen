package handlers

import (
	"SchoolAPI/internal/models"
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
	// Implementation here
}
func UpdateStudent() {
	// Implementation here
}
func DeleteStudent() {
	// Implementation here
}
