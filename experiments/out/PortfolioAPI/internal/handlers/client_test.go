package handlers_test

import (
	"PortfolioAPI/internal/handlers"
	"PortfolioAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateClient(t *testing.T) {
	client := models.Client{
		Name:  "Test Client",
		Email: "test@example.com",
	}
	body, _ := json.Marshal(client)
	req := httptest.NewRequest("POST", "/clients", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.CreateClient(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetClients(t *testing.T) {
	req := httptest.NewRequest("GET", "/clients", nil)
	w := httptest.NewRecorder()
	handlers.GetClients(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
