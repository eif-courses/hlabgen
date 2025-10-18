package handlers

import (
	"ClinicAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateDoctor() {
	var doctor models.Doctor
	if err := json.NewDecoder(r.Body).Decode(&doctor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(doctor)
}

func GetDoctors() {
	// Implementation for getting doctors
}
func UpdateDoctor() {
	// Implementation for updating a doctor
}
func DeleteDoctor() {
	// Implementation for deleting a doctor
}
