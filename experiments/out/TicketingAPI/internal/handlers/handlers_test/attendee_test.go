package handlers_test

import (
	"TicketingAPI/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateAttendee() {
	req, err := http.NewRequest("POST", "/attendees", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateAttendee)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}
