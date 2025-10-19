package handlers_test

import (
	"CourseAPI/internal/handlers"
	"CourseAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateLesson(t *testing.T) {
	lesson := models.Lesson{
		CourseID: 1,
		Title:    "Test Lesson",
		Content:  "Test Content",
	}
	body, _ := json.Marshal(lesson)
	req := httptest.NewRequest("POST", "/lessons", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateLesson(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetLessons(t *testing.T) {
	req := httptest.NewRequest("GET", "/lessons", nil)
	w := httptest.NewRecorder()
	handlers.GetLessons(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
