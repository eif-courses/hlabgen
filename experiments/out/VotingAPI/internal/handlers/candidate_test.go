package handlers_test

import (
	"VotingAPI/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateCandidate() {
	req, err := http.NewRequest("POST", "/candidates", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	handlers.CreateCandidate(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %v", w.Code)
	}
}
