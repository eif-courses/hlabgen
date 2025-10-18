package handlers

import (
	"CourseAPI/internal/models"
	"encoding/json"
	"net/http"
)

// CreateEnrollment handles the creation of a new enrollment.
func CreateEnrollment() {
	var enrollment models.Enrollment
	if err := json.NewDecoder(r.Body).Decode(&enrollment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Logic to save enrollment to database
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(enrollment)
}

// GetEnrollments handles fetching all enrollments.
func GetEnrollments() {
	// Logic to retrieve enrollments from database
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode([]models.Enrollment{})
}
