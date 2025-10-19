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

func TestCreateMilestone(t *testing.T) {
	milestone := models.Milestone{
		Title:     "Test Milestone",
		DueDate:   "2022-12-31",
		ProjectID: 1,
	}
	body, _ := json.Marshal(milestone)
	req := httptest.NewRequest("POST", "/milestones", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.CreateMilestone(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetMilestones(t *testing.T) {
	req := httptest.NewRequest("GET", "/milestones", nil)
	w := httptest.NewRecorder()
	handlers.GetMilestones(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
