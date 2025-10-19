package handlers_test

import (
	"MusicAPI/internal/handlers"
	"MusicAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateArtist(t *testing.T) {
	artist := models.Artist{
		Name: "Test Artist",
	}
	body, _ := json.Marshal(artist)
	req := httptest.NewRequest("POST", "/artists", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.CreateArtist(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetArtists(t *testing.T) {
	req := httptest.NewRequest("GET", "/artists", nil)
	w := httptest.NewRecorder()
	handlers.GetArtists(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
