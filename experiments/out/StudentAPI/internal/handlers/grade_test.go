package handlers_test

import (
	"StudentAPI/internal/handlers"
	"StudentAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateGrade(t *testing.T) {
	grade := models.Grade{
		StudentID: 1,
		CourseID:  1,
		Grade:     90,
	}
	body, _ := json.Marshal(grade)
	req := httptest.NewRequest("POST", "/grades", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateGrade(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
