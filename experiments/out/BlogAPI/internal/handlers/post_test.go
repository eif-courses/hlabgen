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

func TestCreatePost(t *testing.T) {
	post := models.Post{
		Title:   "Test Post",
		Content: "Test Content",
		Author:  "Test Author",
	}
	body, _ := json.Marshal(post)
	req := httptest.NewRequest("POST", "/posts", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreatePost(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetPosts(t *testing.T) {
	req := httptest.NewRequest("GET", "/posts", nil)
	w := httptest.NewRecorder()
	handlers.GetPosts(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
