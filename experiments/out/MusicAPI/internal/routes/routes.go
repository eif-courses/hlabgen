package routes

import (
	"MusicAPI/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	router.POST("/songs", handlers.CreateSong)
	router.GET("/songs/:id", handlers.GetSong)
	router.PUT("/songs/:id", handlers.UpdateSong)
	router.DELETE("/songs/:id", handlers.DeleteSong)
	router.GET("/songs", handlers.GetAllSongs)

	router.POST("/albums", handlers.CreateAlbum)
	router.GET("/albums/:id", handlers.GetAlbum)
	router.PUT("/albums/:id", handlers.UpdateAlbum)
	router.DELETE("/albums/:id", handlers.DeleteAlbum)
	router.GET("/albums", handlers.GetAllAlbums)

	router.POST("/artists", handlers.CreateArtist)
	router.GET("/artists/:id", handlers.GetArtist)
	router.PUT("/artists/:id", handlers.UpdateArtist)
	router.DELETE("/artists/:id", handlers.DeleteArtist)
	router.GET("/artists", handlers.GetAllArtists)

	router.POST("/playlists", handlers.CreatePlaylist)
	router.GET("/playlists/:id", handlers.GetPlaylist)
	router.PUT("/playlists/:id", handlers.UpdatePlaylist)
	router.DELETE("/playlists/:id", handlers.DeletePlaylist)
	router.GET("/playlists", handlers.GetAllPlaylists)

	router.POST("/users", handlers.CreateUser)
	router.GET("/users/:id", handlers.GetUser)
	router.PUT("/users/:id", handlers.UpdateUser)
	router.DELETE("/users/:id", handlers.DeleteUser)
	router.GET("/users", handlers.GetAllUsers)

	router.POST("/likes", handlers.CreateLike)
	router.DELETE("/likes/:id", handlers.DeleteLike)
	router.GET("/users/:id/likes", handlers.GetLikesByUser)
}
