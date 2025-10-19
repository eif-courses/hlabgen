package handlers_test

import (
	"HotelAPI/internal/handlers"
	"HotelAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateHotel(t *testing.T) {
	hotel := models.Hotel{
		Name:    "Test Hotel",
		Address: "Test Address",
		Rating:  5,
	}
	body, _ := json.Marshal(hotel)
	req := httptest.NewRequest("POST", "/hotels", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateHotel(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
