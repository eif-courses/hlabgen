package handlers

import (
    "encoding/json"
    "net/http"
    "HotelAPI/internal/models"
)

func CreateHotel(w http.ResponseWriter, r *http.Request) {
    var hotel models.Hotel
    if err := json.NewDecoder(r.Body).Decode(&hotel); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(hotel)
}

func GetHotels(w http.ResponseWriter, r *http.Request) {
    // Implementation for fetching hotels
}

func UpdateHotel(w http.ResponseWriter, r *http.Request) {
    // Implementation for updating a hotel
}

func DeleteHotel(w http.ResponseWriter, r *http.Request) {
    // Implementation for deleting a hotel
}