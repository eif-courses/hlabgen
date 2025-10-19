package handlers_test

import (
	"VotingAPI/internal/handlers"
	"VotingAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateElection(t *testing.T) {
	election := models.Election{
		Title:    "Test Election",
		Date:     "2022-12-31",
		Location: "Test Location",
	}
	body, _ := json.Marshal(election)
	req := httptest.NewRequest("POST", "/elections", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.CreateElection(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetElections(t *testing.T) {
	req := httptest.NewRequest("GET", "/elections", nil)
	w := httptest.NewRecorder()
	handlers.GetElections(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
