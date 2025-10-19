package handlers_test

import (
	"CarRentalAPI/internal/handlers"
	"CarRentalAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateCustomer(t *testing.T) {
	customer := models.Customer{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password",
	}
	body, _ := json.Marshal(customer)
	req := httptest.NewRequest("POST", "/customers", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateCustomer(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetCustomers(t *testing.T) {
	req := httptest.NewRequest("GET", "/customers", nil)
	w := httptest.NewRecorder()
	handlers.GetCustomers(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
