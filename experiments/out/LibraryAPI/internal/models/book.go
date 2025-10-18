package models

import "time"

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	PublishedAt time.Time `json:"published_at"`
}
