package handlers

import (
	"CourseAPI/internal/models"
	"encoding/json"
	"net/http"
)

var enrollments []models.Enrollment

func CreateEnrollment(w http.ResponseWriter, r *http.Request) {
	var enrollment models.Enrollment
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&enrollment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	enrollments = append(enrollments, enrollment)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(enrollment)
}

func GetEnrollments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(enrollments)
}
func GetEnrollment(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single enrollment
	w.WriteHeader(http.StatusOK)
}
func UpdateEnrollment(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating an enrollment
	w.WriteHeader(http.StatusOK)
}
func DeleteEnrollment(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting an enrollment
	w.WriteHeader(http.StatusNoContent)
}
