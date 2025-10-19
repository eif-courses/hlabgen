package handlers_test

import (
	"FitnessAPI/internal/handlers"
	"FitnessAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateGoal(t *testing.T) {
	goal := models.Goal{
		UserID:   1,
		Target:   5000,
		Achieved: 2500,
	}
	body, _ := json.Marshal(goal)
	req := httptest.NewRequest("POST", "/goals", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateGoal(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
