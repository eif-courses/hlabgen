package handlers_test

import (
	"FoodDeliveryAPI/internal/handlers"
	"FoodDeliveryAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateMenuItem(t *testing.T) {
	menuItem := models.MenuItem{
		Name:     "Test Item",
		Price:    9.99,
		Category: "Test Category",
	}
	body, _ := json.Marshal(menuItem)
	req := httptest.NewRequest("POST", "/menuitems", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateMenuItem(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
