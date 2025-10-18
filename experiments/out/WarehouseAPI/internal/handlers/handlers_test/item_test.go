package handlers_test

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"
    "WarehouseAPI/internal/handlers"
    "WarehouseAPI/internal/models"
)

func TestCreateItem() {
item := models.Item{Name: "Test Item", Quantity: 10, LocationID: 1
}
    body, _ := json.Marshal(item)
    req, err := http.NewRequest("POST", "/items", bytes.NewBuffer(body))
    if err != nil {
        t.Fatal(err)
    }
    w := httptest.NewRecorder()
    handlers.CreateItem(w, req)

    if w.Code != http.StatusCreated {
        t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
    }
}
