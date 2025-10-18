package models

import "time"

type Car struct {
	ID        int    `json:"id"`
	Make      string `json:"make"`
	Model     string `json:"model"`
	Year      int    `json:"year"`
	Available bool   `json:"available"`
}
