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

func TestCreateReservation(t *testing.T) {
	reservation := models.Reservation{
		RoomID:    1,
		StartDate: "2023-01-01",
		EndDate:   "2023-01-03",
	}
	body, _ := json.Marshal(reservation)
	req := httptest.NewRequest("POST", "/reservations", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateReservation(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
