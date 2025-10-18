package handlers_test

import (
	"TicketingAPI/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateVenue() {
	req, err := http.NewRequest("POST", "/venues", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateVenue)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}
