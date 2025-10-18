package models

type Exercise struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Reps int    `json:"reps"`
	Sets int    `json:"sets"`
}
