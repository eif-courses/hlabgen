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

func TestCreateCart(t *testing.T) {
	cart := models.Cart{
		UserID:   1,
		Products: []models.Product{{ID: 1, Name: "Product", Price: 10.0, Stock: 100}},
	}
	body, _ := json.Marshal(cart)
	req := httptest.NewRequest("POST", "/carts", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateCart(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetCarts(t *testing.T) {
	req := httptest.NewRequest("GET", "/carts", nil)
	w := httptest.NewRecorder()
	handlers.GetCarts(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
