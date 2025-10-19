package handlers_test

import (
	"ClinicAPI/internal/handlers"
	"ClinicAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateAppointment(t *testing.T) {
	appointment := models.Appointment{
		DoctorID:  1,
		PatientID: 1,
		DateTime:  "2022-01-01T10:00:00",
	}
	body, _ := json.Marshal(appointment)
	req := httptest.NewRequest("POST", "/appointments", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateAppointment(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetAppointments(t *testing.T) {
	req := httptest.NewRequest("GET", "/appointments", nil)
	w := httptest.NewRecorder()
	handlers.GetAppointments(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
