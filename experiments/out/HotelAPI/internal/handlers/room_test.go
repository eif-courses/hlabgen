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

func TestCreateRoom(t *testing.T) {
	room := models.Room{
		Number:   "101",
		Capacity: 2,
		HotelID:  1,
	}
	body, _ := json.Marshal(room)
	req := httptest.NewRequest("POST", "/rooms", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateRoom(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
