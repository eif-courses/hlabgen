package handlers

import (
	"HotelAPI/internal/models"
	"encoding/json"
	"net/http"
)

var hotels []models.Hotel

func CreateHotel(w http.ResponseWriter, r *http.Request) {
	var hotel models.Hotel
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&hotel); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	hotels = append(hotels, hotel)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(hotel)
}
