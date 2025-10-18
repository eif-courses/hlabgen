package models

import "time"

// Appointment represents an appointment in the clinic.
type Appointment struct {
	ID        int       `json:"id"`
	DoctorID  int       `json:"doctor_id"`
	PatientID int       `json:"patient_id"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
