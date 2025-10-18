package handlers_test

import (
	"VotingAPI/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateElection() {
	req, err := http.NewRequest("POST", "/elections", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	handlers.CreateElection(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %v", w.Code)
	}
}
