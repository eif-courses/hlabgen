package models

import "time"

type Goal struct {
	ID       int       `json:"id"`
	UserID   int       `json:"user_id"`
	Target   string    `json:"target"`
	Deadline time.Time `json:"deadline"`
}
