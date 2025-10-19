package handlers_test

import (
	"TaskManagerAPI/internal/handlers"
	"TaskManagerAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTeam(t *testing.T) {
	team := models.Team{
		Name: "Test Team",
	}
	body, _ := json.Marshal(team)
	req := httptest.NewRequest("POST", "/teams", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateTeam(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
