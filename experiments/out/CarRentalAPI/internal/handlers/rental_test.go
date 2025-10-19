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
		StartDate:  "2022-01-01",
		EndDate:    "2022-01-10",
	}
	body, _ := json.Marshal(rental)
	req := httptest.NewRequest("POST", "/rentals", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateRental(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
