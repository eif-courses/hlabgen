package handlers

import (
	"HotelAPI/internal/models"
	"encoding/json"
	"net/http"
)

var rooms []models.Room

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	var room models.Room
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rooms = append(rooms, room)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(room)
}
