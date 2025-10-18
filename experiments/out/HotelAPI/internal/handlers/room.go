package handlers

import (
    "encoding/json"
    "net/http"
    "HotelAPI/internal/models"
)

func CreateRoom(w http.ResponseWriter, r *http.Request) {
    var room models.Room
    if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(room)
}

func GetRooms(w http.ResponseWriter, r *http.Request) {
    // Implementation for fetching rooms
}

func UpdateRoom(w http.ResponseWriter, r *http.Request) {
    // Implementation for updating a room
}

func DeleteRoom(w http.ResponseWriter, r *http.Request) {
    // Implementation for deleting a room
}