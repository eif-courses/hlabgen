package models

type Workout struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Date     string `json:"date"`
	Duration int    `json:"duration"`
	Calories int    `json:"calories"`
}
