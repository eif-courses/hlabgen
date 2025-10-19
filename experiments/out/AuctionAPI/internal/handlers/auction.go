package handlers

import (
	"AuctionAPI/internal/models"
	"encoding/json"
	"net/http"
)

var auctions []models.Auction

func CreateAuction(w http.ResponseWriter, r *http.Request) {
	var auction models.Auction
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&auction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	auctions = append(auctions, auction)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(auction)
}
