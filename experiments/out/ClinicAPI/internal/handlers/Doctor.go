package handlers

import (
	"ClinicAPI/internal/models"
	"encoding/json"
	"net/http"
)

var doctors []models.Doctor

func CreateDoctor(w http.ResponseWriter, r *http.Request) {
	var doctor models.Doctor
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&doctor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	doctors = append(doctors, doctor)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(doctor)
}

func GetDoctors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(doctors)
}
func GetDoctor(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single doctor
	w.WriteHeader(http.StatusOK)
}
func UpdateDoctor(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a doctor
	w.WriteHeader(http.StatusOK)
}
func DeleteDoctor(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a doctor
	w.WriteHeader(http.StatusNoContent)
}
