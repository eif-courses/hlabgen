package models

type Goal struct {
	ID       int `json:"id"`
	UserID   int `json:"user_id"`
	Target   int `json:"target"`
	Achieved int `json:"achieved"`
}
