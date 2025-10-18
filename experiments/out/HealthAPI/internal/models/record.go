package models

import "time"

// Record represents a medical record in the system.
type Record struct {
	ID        int       `json:"id"`
	PatientID int       `json:"patient_id"`
	DoctorID  int       `json:"doctor_id"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
