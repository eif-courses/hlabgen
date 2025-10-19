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

func TestCreateBooking(t *testing.T) {
	booking := models.Booking{
		TripID: 1,
		Date:   "2022-12-31",
		Status: "Confirmed",
	}
	body, _ := json.Marshal(booking)
	req := httptest.NewRequest("POST", "/bookings", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateBooking(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetBookings(t *testing.T) {
	req := httptest.NewRequest("GET", "/bookings", nil)
	w := httptest.NewRecorder()
	handlers.GetBookings(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
