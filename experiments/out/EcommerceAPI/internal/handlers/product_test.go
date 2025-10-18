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

func TestCreateProduct(t *testing.T) {
	product := models.Product{
		Name:  "Test Product",
		Price: 10.0,
		Stock: 100,
	}
	body, _ := json.Marshal(product)
	req := httptest.NewRequest("POST", "/products", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateProduct(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetProducts(t *testing.T) {
	req := httptest.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	handlers.GetProducts(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
