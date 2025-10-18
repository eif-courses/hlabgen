package handlers_test

import (
    "bytes"
    "encoding/json"
    "net/http/httptest"
    "testing"
    "EcommerceAPI/internal/handlers"
    "EcommerceAPI/internal/models"
)

func TestCreateCart() {
cart := models.Cart{UserID: 1, Items: []models.Item{{Name: "item1", Price: 10}}, Total: 10
}
    body, _ := json.Marshal(cart)
    req := httptest.NewRequest("POST", "/carts", bytes.NewBuffer(body))
    w := httptest.NewRecorder()
    handlers.CreateCart(w, req)
    if w.Code != http.StatusCreated {
        t.Errorf("Expected 201, got %d", w.Code)
    }
}
