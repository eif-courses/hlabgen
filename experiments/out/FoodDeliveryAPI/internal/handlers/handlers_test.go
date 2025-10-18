package handlers_test

import (
	"FoodDeliveryAPI/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateRestaurant() {
	req := httptest.NewRequest("POST", "/restaurants", nil)
	w := httptest.NewRecorder()
	handlers.CreateRestaurant(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("expected status 201, got %v", w.Code)
	}
}

// Additional tests for other handlers would go here.
