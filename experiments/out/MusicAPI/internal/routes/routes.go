package routes

import (
	"MusicAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	r.HandleFunc("/songs", handlers.CreateSong).Methods("POST")
	r.HandleFunc("/songs/{id:[0-9]+}", handlers.GetSong).Methods("GET")
	r.HandleFunc("/songs/{id:[0-9]+}", handlers.UpdateSong).Methods("PUT")
	r.HandleFunc("/songs/{id:[0-9]+}", handlers.DeleteSong).Methods("DELETE")
	// Add routes for albums, artists, playlists, users, likes
}
