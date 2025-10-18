package models

type Goal struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	GoalType string `json:"goal_type"`
	Target   int    `json:"target"`
	Achieved bool   `json:"achieved"`
}
