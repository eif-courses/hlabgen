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

func TestCreateRental(t *testing.T) {
	rental := models.Rental{
		CarID:      1,
		CustomerID: 1,
		StartDate:  "2023-01-01",
		EndDate:    "2023-01-10",
		TotalPrice: 100.0,
	}
	body, _ := json.Marshal(rental)
	req := httptest.NewRequest("POST", "/rentals", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateRental(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetRentals(t *testing.T) {
	req := httptest.NewRequest("GET", "/rentals", nil)
	w := httptest.NewRecorder()
	handlers.GetRentals(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
