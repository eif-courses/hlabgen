package handlers_test

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"
    "WarehouseAPI/internal/handlers"
    "WarehouseAPI/internal/models"
)

func TestCreateShipment() {
shipment := models.Shipment{ItemID: 1, Quantity: 5, Status: "Pending"
}
    body, _ := json.Marshal(shipment)
    req, err := http.NewRequest("POST", "/shipments", bytes.NewBuffer(body))
    if err != nil {
        t.Fatal(err)
    }
    w := httptest.NewRecorder()
    handlers.CreateShipment(w, req)

    if w.Code != http.StatusCreated {
        t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
    }
}
