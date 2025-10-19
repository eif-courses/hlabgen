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

func TestCreatePatient(t *testing.T) {
	patient := models.Patient{
		Name: "Test Patient",
		Age:  30,
	}
	body, _ := json.Marshal(patient)
	req := httptest.NewRequest("POST", "/patients", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreatePatient(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetPatients(t *testing.T) {
	req := httptest.NewRequest("GET", "/patients", nil)
	w := httptest.NewRecorder()
	handlers.GetPatients(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
