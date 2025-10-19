package handlers_test

import (
	"RecipeAPI/internal/handlers"
	"RecipeAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateStep(t *testing.T) {
	step := models.Step{
		Order: 1,
		Desc:  "Test Step",
	}
	body, _ := json.Marshal(step)
	req := httptest.NewRequest("POST", "/steps", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.CreateStep(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetSteps(t *testing.T) {
	req := httptest.NewRequest("GET", "/steps", nil)
	w := httptest.NewRecorder()
	handlers.GetSteps(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
