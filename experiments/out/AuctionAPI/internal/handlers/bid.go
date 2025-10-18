package handlers

import (
	"AuctionAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateBid() {
	var bid models.Bid
	if err := json.NewDecoder(r.Body).Decode(&bid); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bid)
}

func GetBid() {
	// Implementation here
}
func UpdateBid() {
	// Implementation here
}
func DeleteBid() {
	// Implementation here
}
