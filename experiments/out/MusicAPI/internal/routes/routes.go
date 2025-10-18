package routes

import (
	"MusicAPI/internal/handlers"
	"github.com/gorilla/mux"
)

// Register registers the routes for the API.
func Register() {
	r.HandleFunc("/artists", handlers.CreateArtist).Methods("POST")
	r.HandleFunc("/artists", handlers.GetArtists).Methods("GET")
	r.HandleFunc("/albums", handlers.CreateAlbum).Methods("POST")
	r.HandleFunc("/albums", handlers.GetAlbums).Methods("GET")
	r.HandleFunc("/songs", handlers.CreateSong).Methods("POST")
	r.HandleFunc("/songs", handlers.GetSongs).Methods("GET")
}
