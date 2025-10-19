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

func TestCreateExercise(t *testing.T) {
	exercise := models.Exercise{
		Name:     "Push-up",
		Category: "Strength",
	}
	body, _ := json.Marshal(exercise)
	req := httptest.NewRequest("POST", "/exercises", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateExercise(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
