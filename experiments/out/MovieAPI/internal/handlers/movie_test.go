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

func TestCreateMovie(t *testing.T) {
	movie := models.Movie{
		Title:    "Test Movie",
		Director: "Test Director",
		Year:     2021,
	}
	body, _ := json.Marshal(movie)
	req := httptest.NewRequest("POST", "/movies", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateMovie(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetMovies(t *testing.T) {
	req := httptest.NewRequest("GET", "/movies", nil)
	w := httptest.NewRecorder()
	handlers.GetMovies(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
