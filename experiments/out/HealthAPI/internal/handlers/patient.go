package handlers

import (
	"HealthAPI/internal/models"
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
	// Logic to save patient to the database goes here.
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(patient)
}

// GetPatients handles fetching all patients with pagination.
func GetPatients() {
	// Logic to fetch patients from the database goes here.
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode([]models.Patient{})
}
