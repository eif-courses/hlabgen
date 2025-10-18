package handlers_test

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "encoding/json"
    "HotelAPI/internal/models"
    "HotelAPI/internal/handlers"
)

func TestCreateReservation() {
reservation := models.Reservation{RoomID: 1, GuestName: "John Doe", CheckIn: time.Now(), CheckOut: time.Now().Add(24 * time.Hour)
}
    body, _ := json.Marshal(reservation)
    req, err := http.NewRequest("POST", "/reservations", bytes.NewBuffer(body))
    if err != nil {
        t.Fatal(err)
    }
    w := httptest.NewRecorder()
    handlers.CreateReservation(w, req)
    if w.Code != http.StatusCreated {
        t.Errorf("Expected status 201, got %v", w.Code)
    }
}
