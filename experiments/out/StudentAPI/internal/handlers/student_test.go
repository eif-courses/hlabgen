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

func TestCreateStudent(t *testing.T) {
	student := models.Student{
		Name: "Test Student",
		Age:  20,
	}
	body, _ := json.Marshal(student)
	req := httptest.NewRequest("POST", "/students", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateStudent(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
