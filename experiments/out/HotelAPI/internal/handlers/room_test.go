package handlers_test

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "encoding/json"
    "HotelAPI/internal/models"
    "HotelAPI/internal/handlers"
)

func TestCreateRoom() {
room := models.Room{HotelID: 1, RoomType: "Deluxe", Price: 150.00
}
    body, _ := json.Marshal(room)
    req, err := http.NewRequest("POST", "/rooms", bytes.NewBuffer(body))
    if err != nil {
        t.Fatal(err)
    }
    w := httptest.NewRecorder()
    handlers.CreateRoom(w, req)
    if w.Code != http.StatusCreated {
        t.Errorf("Expected status 201, got %v", w.Code)
    }
}
