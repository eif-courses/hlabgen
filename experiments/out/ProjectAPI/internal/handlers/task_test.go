package handlers_test

import (
	"ProjectAPI/internal/handlers"
	"ProjectAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTask(t *testing.T) {
	task := models.Task{
		Title:       "Test Task",
		Description: "Test Description",
		ProjectID:   1,
	}
	body, _ := json.Marshal(task)
	req := httptest.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.CreateTask(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetTasks(t *testing.T) {
	req := httptest.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	handlers.GetTasks(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
