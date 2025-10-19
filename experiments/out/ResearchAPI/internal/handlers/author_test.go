package handlers_test

import (
	"ResearchAPI/internal/handlers"
	"ResearchAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateAuthor(t *testing.T) {
	author := models.Author{
		Name:        "Test Author",
		Affiliation: "Test University",
	}
	body, _ := json.Marshal(author)
	req := httptest.NewRequest("POST", "/authors", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateAuthor(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
