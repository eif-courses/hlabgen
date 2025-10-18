package models

import "time"

type Survey struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type Question struct {
	ID       int    `json:"id"`
	SurveyID int    `json:"survey_id"`
	Text     string `json:"text"`
}
type Response struct {
	ID         int       `json:"id"`
	QuestionID int       `json:"question_id"`
	Answer     string    `json:"answer"`
	CreatedAt  time.Time `json:"created_at"`
}
