package handlers_test

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"
    "FoodDeliveryAPI/internal/models"
    "FoodDeliveryAPI/internal/handlers"
)

func TestCreateRestaurant() {
restaurant := models.Restaurant{Name: "Test Restaurant", Address: "123 Test St"
}
    body, _ := json.Marshal(restaurant)
    req, err := http.NewRequest("POST", "/restaurants", bytes.NewBuffer(body))
    if err != nil {
        t.Fatal(err)
    }
    w := httptest.NewRecorder()
    handlers.CreateRestaurant(w, req)
    if w.Code != http.StatusCreated {
        t.Errorf("Expected status %v, got %v", http.StatusCreated, w.Code)
    }
}
