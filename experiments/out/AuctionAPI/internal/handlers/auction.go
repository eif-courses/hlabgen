package handlers

import (
	"AuctionAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateAuction() {
	var auction models.Auction
	if err := json.NewDecoder(r.Body).Decode(&auction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(auction)
}

func GetAuction() {
	// Implementation here
}
func UpdateAuction() {
	// Implementation here
}
func DeleteAuction() {
	// Implementation here
}
