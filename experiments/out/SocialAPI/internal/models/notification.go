package models

import "time"

type Notification struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
