package handlers

import (
	"AuctionAPI/internal/models"
	"encoding/json"
	"net/http"
)

var bids []models.Bid

func CreateBid(w http.ResponseWriter, r *http.Request) {
	var bid models.Bid
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&bid); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bids = append(bids, bid)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bid)
}

func GetBids(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bids)
}
func GetBid(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single bid
	w.WriteHeader(http.StatusOK)
}
func UpdateBid(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a bid
	w.WriteHeader(http.StatusOK)
}
func DeleteBid(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a bid
	w.WriteHeader(http.StatusNoContent)
}
