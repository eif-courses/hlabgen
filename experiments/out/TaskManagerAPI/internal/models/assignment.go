package models

import "time"

type Assignment struct {
	ID         int       `json:"id"`
	TaskID     int       `json:"task_id"`
	TeamID     int       `json:"team_id"`
	AssignedAt time.Time `json:"assigned_at"`
}
