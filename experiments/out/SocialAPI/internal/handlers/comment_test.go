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

func TestCreateComment(t *testing.T) {
	comment := models.Comment{
		PostID:  1,
		Content: "This is a test comment",
	}
	body, _ := json.Marshal(comment)
	req := httptest.NewRequest("POST", "/comments", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateComment(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
