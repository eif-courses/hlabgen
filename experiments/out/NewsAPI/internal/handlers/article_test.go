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

func TestCreateArticle(t *testing.T) {
	article := models.Article{
		Title:   "Test Article",
		Content: "Test Content",
		Author:  "Test Author",
	}
	body, _ := json.Marshal(article)
	req := httptest.NewRequest("POST", "/articles", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateArticle(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetArticles(t *testing.T) {
	req := httptest.NewRequest("GET", "/articles", nil)
	w := httptest.NewRecorder()
	handlers.GetArticles(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
