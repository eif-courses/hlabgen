package handlers

import (
	"ClinicAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreatePatient() {
	var patient models.Patient
	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(patient)
}

func GetPatients() {
	// Implementation for getting patients
}
func UpdatePatient() {
	// Implementation for updating a patient
}
func DeletePatient() {
	// Implementation for deleting a patient
}
