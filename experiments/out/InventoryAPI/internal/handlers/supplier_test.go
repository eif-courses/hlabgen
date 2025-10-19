package handlers_test

import (
	"InventoryAPI/internal/handlers"
	"InventoryAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateSupplier(t *testing.T) {
	supplier := models.Supplier{
		Name: "Test Supplier",
	}
	body, _ := json.Marshal(supplier)
	req := httptest.NewRequest("POST", "/suppliers", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateSupplier(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetSuppliers(t *testing.T) {
	req := httptest.NewRequest("GET", "/suppliers", nil)
	w := httptest.NewRecorder()
	handlers.GetSuppliers(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
