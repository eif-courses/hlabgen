package handlers

import (
	"ClinicAPI/internal/models"
	"encoding/json"
	"net/http"
)

// CreatePatient handles the creation of a new patient.
func CreatePatient() {
	var patient models.Patient
	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(patient)
}

// GetPatients handles retrieving all patients.
func GetPatients() {
	// Implementation here
}

// UpdatePatient handles updating an existing patient.
func UpdatePatient() {
	// Implementation here
}

// DeletePatient handles deleting a patient.
func DeletePatient() {
	// Implementation here
}
