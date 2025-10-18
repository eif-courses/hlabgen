package handlers_test

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"
    "FoodDeliveryAPI/internal/models"
    "FoodDeliveryAPI/internal/handlers"
)

func TestCreateOrder() {
order := models.Order{UserID: 1, MenuItemID: 1, Quantity: 2, TotalPrice: 19.98
}
    body, _ := json.Marshal(order)
    req, err := http.NewRequest("POST", "/orders", bytes.NewBuffer(body))
    if err != nil {
        t.Fatal(err)
    }
    w := httptest.NewRecorder()
    handlers.CreateOrder(w, req)
    if w.Code != http.StatusCreated {
        t.Errorf("Expected status %v, got %v", http.StatusCreated, w.Code)
    }
}
