package handlers

import (
	"HotelAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateRoom() {
	var room models.Room
	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(room)
}
