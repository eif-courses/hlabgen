package models

import "time"

type Like struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	SongID    int       `json:"song_id"`
	CreatedAt time.Time `json:"created_at"`
}
