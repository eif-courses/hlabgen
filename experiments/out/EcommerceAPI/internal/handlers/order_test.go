package handlers_test

import (
    "bytes"
    "encoding/json"
    "net/http/httptest"
    "testing"
    "EcommerceAPI/internal/handlers"
    "EcommerceAPI/internal/models"
)

func TestCreateOrder() {
order := models.Order{UserID: 1, Products: []models.Product{{ID: 1, Name: "Test Product", Price: 10.0, Stock: 100}}, Total: 10
}
    body, _ := json.Marshal(order)
    req := httptest.NewRequest("POST", "/orders", bytes.NewBuffer(body))
    w := httptest.NewRecorder()
    handlers.CreateOrder(w, req)
    if w.Code != http.StatusCreated {
        t.Errorf("Expected 201, got %d", w.Code)
    }
}
