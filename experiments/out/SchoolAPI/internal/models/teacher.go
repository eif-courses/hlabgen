package models

import "time"

type Teacher struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Subject   string    `json:"subject"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
