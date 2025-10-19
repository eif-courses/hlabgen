package handlers_test

import (
	"FitnessAPI/internal/handlers"
	"FitnessAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {
	user := models.User{
		Username: "TestUser",
		Email:    "test@example.com",
		Password: "password",
	}
	body, _ := json.Marshal(user)
	req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateUser(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
