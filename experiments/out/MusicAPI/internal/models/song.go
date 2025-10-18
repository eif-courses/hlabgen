package models

type Song struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	AlbumID int    `json:"album_id"`
}
