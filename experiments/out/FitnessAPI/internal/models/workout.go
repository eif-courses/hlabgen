package models

type Workout struct {
	ID        int        `json:"id"`
	UserID    int        `json:"user_id"`
	Date      time.Time  `json:"date"`
	Exercises []Exercise `json:"exercises"`
}
