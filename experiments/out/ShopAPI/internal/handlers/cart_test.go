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

func TestCreateCart(t *testing.T) {
	cart := models.Cart{
		Customer: models.Customer{ID: 1, Name: "Test Customer", Email: "test@example.com"},
		Items:    []models.Product{{ID: 1, Name: "Test Product", Price: 9.99, Stock: 100}},
	}
	body, _ := json.Marshal(cart)
	req := httptest.NewRequest("POST", "/carts", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateCart(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
