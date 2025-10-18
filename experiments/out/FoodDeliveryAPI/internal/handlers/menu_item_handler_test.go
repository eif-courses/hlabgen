package handlers_test

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"
    "FoodDeliveryAPI/internal/models"
    "FoodDeliveryAPI/internal/handlers"
)

func TestCreateMenuItem() {
menuItem := models.MenuItem{Name: "Test Menu Item", Price: 9.99
}
    body, _ := json.Marshal(menuItem)
    req, err := http.NewRequest("POST", "/menuitems", bytes.NewBuffer(body))
    if err != nil {
        t.Fatal(err)
    }
    w := httptest.NewRecorder()
    handlers.CreateMenuItem(w, req)
    if w.Code != http.StatusCreated {
        t.Errorf("Expected status %v, got %v", http.StatusCreated, w.Code)
    }
}
