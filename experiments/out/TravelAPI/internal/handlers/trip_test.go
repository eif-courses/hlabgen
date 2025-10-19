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

func TestCreateTrip(t *testing.T) {
	trip := models.Trip{
		Name:        "Test Trip",
		Description: "Test Description",
		Price:       100.0,
	}
	body, _ := json.Marshal(trip)
	req := httptest.NewRequest("POST", "/trips", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateTrip(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetTrips(t *testing.T) {
	req := httptest.NewRequest("GET", "/trips", nil)
	w := httptest.NewRecorder()
	handlers.GetTrips(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
