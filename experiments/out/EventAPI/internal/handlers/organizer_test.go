package handlers_test

import (
	"EventAPI/internal/handlers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateOrganizer(t *testing.T) {
	body := strings.NewReader("{\"id\":1}")
	req, _ := http.NewRequest("POST", "/organizers", body)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handlers.CreateOrganizer(rr, req)
	if rr.Code != http.StatusCreated && rr.Code != http.StatusOK && rr.Code != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v", rr.Code)
	}
}

func TestGetOrganizer(t *testing.T) {
	body := strings.NewReader("{}")
	req, _ := http.NewRequest("GET", "/organizers", body)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handlers.GetOrganizer(rr, req)
	if rr.Code != http.StatusCreated && rr.Code != http.StatusOK && rr.Code != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v", rr.Code)
	}
}
