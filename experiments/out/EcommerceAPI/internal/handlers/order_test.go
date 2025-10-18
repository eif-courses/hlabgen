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

func TestCreateOrder() {
	order := models.Order{
		UserID: 1,
		Products: []models.Product{
			{ID: 1, Name: "Test Product", Price: 10.0, Stock: 100},
		},
		Total: 10.0,
	}
	body, _ := json.Marshal(order)
	req := httptest.NewRequest("POST", "/orders", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateOrder(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetOrders() {
	req := httptest.NewRequest("GET", "/orders", nil)
	w := httptest.NewRecorder()
	handlers.GetOrders(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
