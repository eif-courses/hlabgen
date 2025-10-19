package handlers_test

import (
	"ShopAPI/internal/handlers"
	"ShopAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	order := models.Order{
		Customer: models.Customer{ID: 1, Name: "Test Customer", Email: "test@example.com"},
		Products: []models.Product{{ID: 1, Name: "Test Product", Price: 9.99, Stock: 100}},
		Total:    9.99,
	}
	body, _ := json.Marshal(order)
	req := httptest.NewRequest("POST", "/orders", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateOrder(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
