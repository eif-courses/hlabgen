package handlers

import (
	"ClinicAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateAppointment() {
	var appointment models.Appointment
	if err := json.NewDecoder(r.Body).Decode(&appointment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(appointment)
}

func GetAppointments() {
	// Implementation for getting appointments
}
func UpdateAppointment() {
	// Implementation for updating an appointment
}
func DeleteAppointment() {
	// Implementation for deleting an appointment
}
