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

func TestCreatePaper(t *testing.T) {
	paper := models.Paper{
		Title:   "Test Paper",
		Author:  "Test Author",
		Content: "Lorem ipsum",
	}
	body, _ := json.Marshal(paper)
	req := httptest.NewRequest("POST", "/papers", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreatePaper(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
