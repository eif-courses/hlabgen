package handlers

import (
	"SchoolAPI/internal/models"
	"encoding/json"
	"net/http"
)

var teachers []models.Teacher

func CreateTeacher(w http.ResponseWriter, r *http.Request) {
	var teacher models.Teacher
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&teacher); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	teachers = append(teachers, teacher)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(teacher)
}

func GetTeachers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teachers)
}
func GetTeacher(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single teacher
	w.WriteHeader(http.StatusOK)
}
func UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a teacher
	w.WriteHeader(http.StatusOK)
}
func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a teacher
	w.WriteHeader(http.StatusNoContent)
}
