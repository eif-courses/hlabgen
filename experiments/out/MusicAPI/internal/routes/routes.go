package routes

import (
	"MusicAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register() {
	r.HandleFunc("/artists", handlers.CreateArtist).Methods("POST")
	r.HandleFunc("/artists", handlers.GetArtists).Methods("GET")
	r.HandleFunc("/artists/{id}", handlers.UpdateArtist).Methods("PUT")
	r.HandleFunc("/artists/{id}", handlers.DeleteArtist).Methods("DELETE")
	r.HandleFunc("/albums", handlers.CreateAlbum).Methods("POST")
	r.HandleFunc("/albums", handlers.GetAlbums).Methods("GET")
	r.HandleFunc("/albums/{id}", handlers.UpdateAlbum).Methods("PUT")
	r.HandleFunc("/albums/{id}", handlers.DeleteAlbum).Methods("DELETE")
	r.HandleFunc("/songs", handlers.CreateSong).Methods("POST")
	r.HandleFunc("/songs", handlers.GetSongs).Methods("GET")
	r.HandleFunc("/songs/{id}", handlers.UpdateSong).Methods("PUT")
	r.HandleFunc("/songs/{id}", handlers.DeleteSong).Methods("DELETE")
}
