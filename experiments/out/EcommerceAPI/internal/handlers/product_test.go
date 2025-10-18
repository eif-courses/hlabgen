package handlers_test

import (
    "bytes"
    "encoding/json"
    "net/http/httptest"
    "testing"
    "EcommerceAPI/internal/handlers"
    "EcommerceAPI/internal/models"
)

func TestCreateProduct() {
product := models.Product{Name: "Test Product", Price: 10.0, Stock: 100
}
    body, _ := json.Marshal(product)
    req := httptest.NewRequest("POST", "/products", bytes.NewBuffer(body))
    w := httptest.NewRecorder()
    handlers.CreateProduct(w, req)
    if w.Code != http.StatusCreated {
        t.Errorf("Expected 201, got %d", w.Code)
    }
}
