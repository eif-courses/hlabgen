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

func TestCreateDoctor(t *testing.T) {
	doctor := models.Doctor{
		Name:      "Test Doctor",
		Specialty: "Cardiology",
	}
	body, _ := json.Marshal(doctor)
	req := httptest.NewRequest("POST", "/doctors", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateDoctor(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetDoctors(t *testing.T) {
	req := httptest.NewRequest("GET", "/doctors", nil)
	w := httptest.NewRecorder()
	handlers.GetDoctors(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
