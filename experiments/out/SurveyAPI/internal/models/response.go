package models

import "time"

type Response struct {
	ID         int       `json:"id"`
	QuestionID int       `json:"question_id"`
	Answer     string    `json:"answer"`
	CreatedAt  time.Time `json:"created_at"`
}
