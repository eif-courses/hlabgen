package handlers_test

import (
	"TravelAPI/internal/handlers"
	"TravelAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateDestination(t *testing.T) {
	destination := models.Destination{
		Name:        "Test Destination",
		Description: "Test Description",
		Country:     "Test Country",
	}
	body, _ := json.Marshal(destination)
	req := httptest.NewRequest("POST", "/destinations", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateDestination(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetDestinations(t *testing.T) {
	req := httptest.NewRequest("GET", "/destinations", nil)
	w := httptest.NewRecorder()
	handlers.GetDestinations(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
