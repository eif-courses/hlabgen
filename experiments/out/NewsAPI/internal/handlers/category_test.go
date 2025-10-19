package handlers_test

import (
	"NewsAPI/internal/handlers"
	"NewsAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateCategory(t *testing.T) {
	category := models.Category{
		Name: "Test Category",
	}
	body, _ := json.Marshal(category)
	req := httptest.NewRequest("POST", "/categories", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateCategory(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetCategories(t *testing.T) {
	req := httptest.NewRequest("GET", "/categories", nil)
	w := httptest.NewRecorder()
	handlers.GetCategories(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
