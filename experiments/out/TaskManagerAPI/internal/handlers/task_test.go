package handlers_test

import (
	"TaskManagerAPI/internal/handlers"
	"TaskManagerAPI/internal/models"
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
		Completed:   false,
	}
	body, _ := json.Marshal(task)
	req := httptest.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateTask(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
