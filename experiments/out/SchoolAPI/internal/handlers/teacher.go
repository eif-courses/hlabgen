package handlers

import (
	"SchoolAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateTeacher() {
	var teacher models.Teacher
	if err := json.NewDecoder(r.Body).Decode(&teacher); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(teacher)
}

func GetTeachers() {
	// Implementation here
}
func UpdateTeacher() {
	// Implementation here
}
func DeleteTeacher() {
	// Implementation here
}
