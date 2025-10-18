package models

import "time"

type Invoice struct {
	ID        int       `json:"id"`
	ProjectID int       `json:"project_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
