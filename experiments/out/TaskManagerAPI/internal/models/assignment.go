package models

type Assignment struct {
	ID     int `json:"id"`
	TaskID int `json:"task_id"`
	TeamID int `json:"team_id"`
}
