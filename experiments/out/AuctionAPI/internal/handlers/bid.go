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
