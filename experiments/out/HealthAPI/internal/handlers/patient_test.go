package handlers_test

import (
	"HealthAPI/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePatient() {
	req, err := http.NewRequest("POST", "/patients", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	handlers.CreatePatient(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, w.Code)
	}
}
