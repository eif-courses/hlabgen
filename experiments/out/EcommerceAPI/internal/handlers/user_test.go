package handlers_test

import (
	"EcommerceAPI/internal/handlers"
	"EcommerceAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser() {
	user := models.User{
		Username: "testuser",
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

func TestGetUsers() {
	req := httptest.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	handlers.GetUsers(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
