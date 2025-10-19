package handlers_test

import (
	"WarehouseAPI/internal/handlers"
	"WarehouseAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateLocation(t *testing.T) {
	location := models.Location{
		Name: "Test Location",
	}
	body, _ := json.Marshal(location)
	req := httptest.NewRequest("POST", "/locations", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.CreateLocation(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetLocations(t *testing.T) {
	req := httptest.NewRequest("GET", "/locations", nil)
	w := httptest.NewRecorder()
	handlers.GetLocations(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
