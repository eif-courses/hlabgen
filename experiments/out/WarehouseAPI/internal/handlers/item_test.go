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

func TestCreateItem(t *testing.T) {
	item := models.Item{
		Name:  "Test Item",
		Price: 20.0,
	}
	body, _ := json.Marshal(item)
	req := httptest.NewRequest("POST", "/items", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.CreateItem(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetItems(t *testing.T) {
	req := httptest.NewRequest("GET", "/items", nil)
	w := httptest.NewRecorder()
	handlers.GetItems(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
