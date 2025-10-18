package handlers

import (
    "encoding/json"
    "net/http"
    "HotelAPI/internal/models"
)

func CreateBooking(w http.ResponseWriter, r *http.Request) {
    var booking models.Booking
    if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(booking)
}

func GetBookings(w http.ResponseWriter, r *http.Request) {
    // Implementation for fetching bookings
}

func UpdateBooking(w http.ResponseWriter, r *http.Request) {
    // Implementation for updating a booking
}

func DeleteBooking(w http.ResponseWriter, r *http.Request) {
    // Implementation for deleting a booking
}