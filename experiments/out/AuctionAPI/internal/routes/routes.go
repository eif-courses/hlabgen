package routes

import (
	"AuctionAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register() {
	r.HandleFunc("/auctions", handlers.CreateAuction).Methods("POST")
	r.HandleFunc("/auctions/{id:[0-9]+}", handlers.GetAuction).Methods("GET")
	r.HandleFunc("/auctions/{id:[0-9]+}", handlers.UpdateAuction).Methods("PUT")
	r.HandleFunc("/auctions/{id:[0-9]+}", handlers.DeleteAuction).Methods("DELETE")
	r.HandleFunc("/bids", handlers.CreateBid).Methods("POST")
	r.HandleFunc("/bids/{id:[0-9]+}", handlers.GetBid).Methods("GET")
	r.HandleFunc("/bids/{id:[0-9]+}", handlers.UpdateBid).Methods("PUT")
	r.HandleFunc("/bids/{id:[0-9]+}", handlers.DeleteBid).Methods("DELETE")
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")
}
