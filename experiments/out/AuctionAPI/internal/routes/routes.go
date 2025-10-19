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
	// Bid routes
	r.HandleFunc("/bids", handlers.CreateBid).Methods("POST")
	r.HandleFunc("/bids", handlers.GetBids).Methods("GET")
	r.HandleFunc("/bids/{id}", handlers.GetBid).Methods("GET")
	r.HandleFunc("/bids/{id}", handlers.UpdateBid).Methods("PUT")
	r.HandleFunc("/bids/{id}", handlers.DeleteBid).Methods("DELETE")
	// User routes
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
}
