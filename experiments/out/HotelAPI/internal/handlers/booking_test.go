package handlers_test

import (
	"HotelAPI/internal/handlers"
	"HotelAPI/internal/routes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateBooking(t *testing.T) {
	body := strings.NewReader("{\"id\":1}")
	req, _ := http.NewRequest("POST", "/bookings", body)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handlers.CreateBooking(rr, req)
	if rr.Code != http.StatusCreated && rr.Code != http.StatusOK && rr.Code != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v", rr.Code)
	}
}

func TestGetBookings(t *testing.T) {
	body := strings.NewReader("{}")
	req, _ := http.NewRequest("GET", "/bookings", body)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handlers.GetBookings(rr, req)
	if rr.Code != http.StatusCreated && rr.Code != http.StatusOK && rr.Code != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v", rr.Code)
	}
}
