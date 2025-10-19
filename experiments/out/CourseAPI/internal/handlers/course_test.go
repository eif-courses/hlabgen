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

func TestCreateCourse(t *testing.T) {
	course := models.Course{
		Title: "Test Course",
	}
	body, _ := json.Marshal(course)
	req := httptest.NewRequest("POST", "/courses", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.CreateCourse(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetCourses(t *testing.T) {
	req := httptest.NewRequest("GET", "/courses", nil)
	w := httptest.NewRecorder()
	handlers.GetCourses(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
