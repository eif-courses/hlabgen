package handlers

import (
	"ClinicAPI/internal/models"
	"encoding/json"
	"net/http"
)

var appointments []models.Appointment

func CreateAppointment(w http.ResponseWriter, r *http.Request) {
	var appointment models.Appointment
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&appointment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	appointments = append(appointments, appointment)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(appointment)
}

func GetAppointments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appointments)
}
func GetAppointment(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single appointment
	w.WriteHeader(http.StatusOK)
}
func UpdateAppointment(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating an appointment
	w.WriteHeader(http.StatusOK)
}
func DeleteAppointment(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting an appointment
	w.WriteHeader(http.StatusNoContent)
}
