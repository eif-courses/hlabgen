package routes

import (
	"AuctionAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	// Auction routes
	r.HandleFunc("/auctions", handlers.CreateAuction).Methods("POST")
	r.HandleFunc("/auctions", handlers.GetAuctions).Methods("GET")
	r.HandleFunc("/auctions/{id}", handlers.GetAuction).Methods("GET")
	r.HandleFunc("/auctions/{id}", handlers.UpdateAuction).Methods("PUT")
	r.HandleFunc("/auctions/{id}", handlers.DeleteAuction).Methods("DELETE")
}
