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

func TestCreateWorkout(t *testing.T) {
	workout := models.Workout{
		UserID:   1,
		Date:     "2022-01-01",
		Duration: 60,
		Calories: 300,
	}
	body, _ := json.Marshal(workout)
	req := httptest.NewRequest("POST", "/workouts", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateWorkout(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
