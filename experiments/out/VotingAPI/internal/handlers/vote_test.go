package handlers_test

import (
	"VotingAPI/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateVote() {
	req, err := http.NewRequest("POST", "/votes", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	handlers.CreateVote(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %v", w.Code)
	}
}
