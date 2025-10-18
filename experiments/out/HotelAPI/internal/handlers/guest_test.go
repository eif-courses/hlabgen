package handlers_test

import (
	"HotelAPI/internal/handlers"
	"HotelAPI/internal/routes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateGuest(t *testing.T) {
	body := strings.NewReader("{\"id\":1}")
	req, _ := http.NewRequest("POST", "/guests", body)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handlers.CreateGuest(rr, req)
	if rr.Code != http.StatusCreated && rr.Code != http.StatusOK && rr.Code != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v", rr.Code)
	}
}

func TestGetGuests(t *testing.T) {
	body := strings.NewReader("{}")
	req, _ := http.NewRequest("GET", "/guests", body)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handlers.GetGuests(rr, req)
	if rr.Code != http.StatusCreated && rr.Code != http.StatusOK && rr.Code != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v", rr.Code)
	}
}
