package models

type Like struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	SongID int `json:"song_id"`
}
