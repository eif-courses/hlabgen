package handlers_test

import (
	"EventAPI/internal/handlers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateAttendee(t *testing.T) {
	body := strings.NewReader("{\"id\":1}")
	req, _ := http.NewRequest("POST", "/attendees", body)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handlers.CreateAttendee(rr, req)
	if rr.Code != http.StatusCreated && rr.Code != http.StatusOK && rr.Code != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v", rr.Code)
	}
}

func TestGetAttendee(t *testing.T) {
	body := strings.NewReader("{}")
	req, _ := http.NewRequest("GET", "/attendees", body)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handlers.GetAttendee(rr, req)
	if rr.Code != http.StatusCreated && rr.Code != http.StatusOK && rr.Code != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v", rr.Code)
	}
}
