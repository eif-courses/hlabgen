package models

type Progress struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	WorkoutID int       `json:"workout_id"`
	Date      time.Time `json:"date"`
	Notes     string    `json:"notes"`
}
