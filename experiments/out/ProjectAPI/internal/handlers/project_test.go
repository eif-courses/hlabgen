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

func TestCreateProject(t *testing.T) {
	project := models.Project{
		Name:        "Test Project",
		Description: "Test Description",
	}
	body, _ := json.Marshal(project)
	req := httptest.NewRequest("POST", "/projects", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.CreateProject(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetProjects(t *testing.T) {
	req := httptest.NewRequest("GET", "/projects", nil)
	w := httptest.NewRecorder()
	handlers.GetProjects(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
