package routes

import (
	"MusicAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	// Artist routes
	r.HandleFunc("/artists", handlers.CreateArtist).Methods("POST")
	r.HandleFunc("/artists", handlers.GetArtists).Methods("GET")
	r.HandleFunc("/artists/{id}", handlers.GetArtist).Methods("GET")
	r.HandleFunc("/artists/{id}", handlers.UpdateArtist).Methods("PUT")
	r.HandleFunc("/artists/{id}", handlers.DeleteArtist).Methods("DELETE")
}
