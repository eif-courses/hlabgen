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

func TestCreateRestaurant(t *testing.T) {
	restaurant := models.Restaurant{
		Name:    "Test Restaurant",
		Address: "Test Address",
		Rating:  4.5,
	}
	body, _ := json.Marshal(restaurant)
	req := httptest.NewRequest("POST", "/restaurants", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateRestaurant(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
