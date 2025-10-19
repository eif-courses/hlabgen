package models

type Exercise struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}
