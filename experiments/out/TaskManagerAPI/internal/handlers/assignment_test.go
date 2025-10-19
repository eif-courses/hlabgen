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

func TestCreateAssignment(t *testing.T) {
	assignment := models.Assignment{
		TaskID: 1,
		TeamID: 1,
	}
	body, _ := json.Marshal(assignment)
	req := httptest.NewRequest("POST", "/assignments", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateAssignment(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
