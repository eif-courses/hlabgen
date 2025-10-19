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

func GetAuctions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(auctions)
}
func GetAuction(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single auction
	w.WriteHeader(http.StatusOK)
}
func UpdateAuction(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating an auction
	w.WriteHeader(http.StatusOK)
}
func DeleteAuction(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting an auction
	w.WriteHeader(http.StatusNoContent)
}
