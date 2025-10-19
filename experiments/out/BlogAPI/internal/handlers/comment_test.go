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

func TestCreateComment(t *testing.T) {
	comment := models.Comment{
		PostID:  1,
		Content: "Test Comment",
		Author:  "Test Author",
	}
	body, _ := json.Marshal(comment)
	req := httptest.NewRequest("POST", "/comments", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateComment(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetComments(t *testing.T) {
	req := httptest.NewRequest("GET", "/comments", nil)
	w := httptest.NewRecorder()
	handlers.GetComments(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
