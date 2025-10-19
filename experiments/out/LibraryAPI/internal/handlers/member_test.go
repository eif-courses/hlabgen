package handlers_test

import (
	"LibraryAPI/internal/handlers"
	"LibraryAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateMember(t *testing.T) {
	member := models.Member{
		Name:  "Test Member",
		Email: "test@example.com",
	}
	body, _ := json.Marshal(member)
	req := httptest.NewRequest("POST", "/members", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.CreateMember(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetMembers(t *testing.T) {
	req := httptest.NewRequest("GET", "/members", nil)
	w := httptest.NewRecorder()
	handlers.GetMembers(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
