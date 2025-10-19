package handlers_test

import (
	"MovieAPI/internal/handlers"
	"MovieAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateDirector(t *testing.T) {
	director := models.Director{
		Name: "Test Director",
		Age:  40,
	}
	body, _ := json.Marshal(director)
	req := httptest.NewRequest("POST", "/directors", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateDirector(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetDirectors(t *testing.T) {
	req := httptest.NewRequest("GET", "/directors", nil)
	w := httptest.NewRecorder()
	handlers.GetDirectors(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
