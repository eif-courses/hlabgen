package handlers_test

import (
	"SchoolAPI/internal/handlers"
	"SchoolAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTeacher(t *testing.T) {
	teacher := models.Teacher{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	}
	body, _ := json.Marshal(teacher)
	req := httptest.NewRequest("POST", "/teachers", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.CreateTeacher(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetTeachers(t *testing.T) {
	req := httptest.NewRequest("GET", "/teachers", nil)
	w := httptest.NewRecorder()
	handlers.GetTeachers(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
