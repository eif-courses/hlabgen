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

func TestCreateCourse(t *testing.T) {
	course := models.Course{
		Name: "Test Course",
		Code: "COURSE101",
	}
	body, _ := json.Marshal(course)
	req := httptest.NewRequest("POST", "/courses", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateCourse(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
