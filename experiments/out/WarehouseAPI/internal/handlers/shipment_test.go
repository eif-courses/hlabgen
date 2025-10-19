package handlers_test

import (
	"WarehouseAPI/internal/handlers"
	"WarehouseAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateShipment(t *testing.T) {
	shipment := models.Shipment{
		Origin:      "Origin",
		Destination: "Destination",
	}
	body, _ := json.Marshal(shipment)
	req := httptest.NewRequest("POST", "/shipments", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.CreateShipment(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetShipments(t *testing.T) {
	req := httptest.NewRequest("GET", "/shipments", nil)
	w := httptest.NewRecorder()
	handlers.GetShipments(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
