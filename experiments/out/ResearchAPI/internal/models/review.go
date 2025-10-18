package models

import "time"

type Review struct {
	ID         int       `json:"id"`
	PaperID    int       `json:"paper_id"`
	ReviewerID int       `json:"reviewer_id"`
	Comments   string    `json:"comments"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
