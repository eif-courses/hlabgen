package handlers

import (
	"ClinicAPI/internal/models"
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
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(doctor)
}

// GetDoctors handles retrieving all doctors.
func GetDoctors() {
	// Implementation here
}

// UpdateDoctor handles updating an existing doctor.
func UpdateDoctor() {
	// Implementation here
}

// DeleteDoctor handles deleting a doctor.
func DeleteDoctor() {
	// Implementation here
}
