package models

type Survey struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Status  string `json:"status"`
	Created string `json:"created"`
}
