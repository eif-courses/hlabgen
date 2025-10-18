package handlers

import (
	"HotelAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateHotel() {
	var hotel models.Hotel
	if err := json.NewDecoder(r.Body).Decode(&hotel); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(hotel)
}

func GetHotels() {
	// Implementation for fetching hotels
}
func UpdateHotel() {
	// Implementation for updating a hotel
}
func DeleteHotel() {
	// Implementation for deleting a hotel
}
