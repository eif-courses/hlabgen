package handlers_test

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "encoding/json"
    "HotelAPI/internal/models"
    "HotelAPI/internal/handlers"
)

func TestCreateHotel() {
hotel := models.Hotel{Name: "Test Hotel", Location: "Test Location"
}
    body, _ := json.Marshal(hotel)
    req, err := http.NewRequest("POST", "/hotels", bytes.NewBuffer(body))
    if err != nil {
        t.Fatal(err)
    }
    w := httptest.NewRecorder()
    handlers.CreateHotel(w, req)
    if w.Code != http.StatusCreated {
        t.Errorf("Expected status 201, got %v", w.Code)
    }
}
