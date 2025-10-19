package handlers

import (
	"ClinicAPI/internal/models"
	"encoding/json"
	"net/http"
)

var patients []models.Patient

func CreatePatient(w http.ResponseWriter, r *http.Request) {
	var patient models.Patient
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	patients = append(patients, patient)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(patient)
}

func GetPatients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patients)
}
func GetPatient(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single patient
	w.WriteHeader(http.StatusOK)
}
func UpdatePatient(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a patient
	w.WriteHeader(http.StatusOK)
}
func DeletePatient(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a patient
	w.WriteHeader(http.StatusNoContent)
}
