package models

type Response struct {
	ID         int    `json:"id"`
	QuestionID int    `json:"question_id"`
	Answer     string `json:"answer"`
	UserID     int    `json:"user_id"`
}
