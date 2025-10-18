package handlers_test

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"
    "WarehouseAPI/internal/handlers"
    "WarehouseAPI/internal/models"
)

func TestCreateLocation() {
location := models.Location{Name: "Test Location"
}
    body, _ := json.Marshal(location)
    req, err := http.NewRequest("POST", "/locations", bytes.NewBuffer(body))
    if err != nil {
        t.Fatal(err)
    }
    w := httptest.NewRecorder()
    handlers.CreateLocation(w, req)
    if w.Code != http.StatusCreated {
        t.Errorf("Expected status 201, got %v", w.Code)
    }
}
