package handlers

import (
	"HealthAPI/internal/models"
	"encoding/json"
	"net/http"
)

// CreateDoctor handles the creation of a new doctor.
func CreateDoctor() {
	var doctor models.Doctor
	if err := json.NewDecoder(r.Body).Decode(&doctor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Logic to save doctor to the database goes here.
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(doctor)
}

// GetDoctors handles fetching all doctors with pagination.
func GetDoctors() {
	// Logic to fetch doctors from the database goes here.
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode([]models.Doctor{})
}
