package models

type Exercise struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Sets int    `json:"sets"`
	Reps int    `json:"reps"`
}
