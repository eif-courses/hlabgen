package models

type Appointment struct {
	ID        int    `json:"id"`
	DoctorID  int    `json:"doctor_id"`
	PatientID int    `json:"patient_id"`
	DateTime  string `json:"date_time"`
}
