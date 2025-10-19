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

func TestCreateEnrollment(t *testing.T) {
	enrollment := models.Enrollment{
		CourseID: 1,
		UserID:   1,
	}
	body, _ := json.Marshal(enrollment)
	req := httptest.NewRequest("POST", "/enrollments", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateEnrollment(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetEnrollments(t *testing.T) {
	req := httptest.NewRequest("GET", "/enrollments", nil)
	w := httptest.NewRecorder()
	handlers.GetEnrollments(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
