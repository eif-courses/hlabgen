package handlers_test

import (
	"BlogAPI/internal/handlers"
	"BlogAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateAuthor(t *testing.T) {
	author := models.Author{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password",
	}
	body, _ := json.Marshal(author)
	req := httptest.NewRequest("POST", "/authors", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateAuthor(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetAuthors(t *testing.T) {
	req := httptest.NewRequest("GET", "/authors", nil)
	w := httptest.NewRecorder()
	handlers.GetAuthors(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
