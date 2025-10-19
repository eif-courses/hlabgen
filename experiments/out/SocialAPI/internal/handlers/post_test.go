package handlers_test

import (
	"SocialAPI/internal/handlers"
	"SocialAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePost(t *testing.T) {
	post := models.Post{
		Title:   "Test Post",
		Content: "This is a test post",
	}
	body, _ := json.Marshal(post)
	req := httptest.NewRequest("POST", "/posts", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreatePost(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
