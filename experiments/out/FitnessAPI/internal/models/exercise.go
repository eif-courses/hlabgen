package models

type Exercise struct {
	ID        int    `json:"id"`
	WorkoutID int    `json:"workout_id"`
	Name      string `json:"name"`
	Reps      int    `json:"reps"`
	Sets      int    `json:"sets"`
}
