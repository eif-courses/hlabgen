package handlers

import (
    "encoding/json"
    "net/http"
    "HotelAPI/internal/models"
)

func CreateGuest(w http.ResponseWriter, r *http.Request) {
    var guest models.Guest
    if err := json.NewDecoder(r.Body).Decode(&guest); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(guest)
}

func GetGuests(w http.ResponseWriter, r *http.Request) {
    // Implementation for fetching guests
}

func UpdateGuest(w http.ResponseWriter, r *http.Request) {
    // Implementation for updating a guest
}

func DeleteGuest(w http.ResponseWriter, r *http.Request) {
    // Implementation for deleting a guest
}