package models

import "time"

type Question struct {
	ID        int       `json:"id"`
	SurveyID  int       `json:"survey_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
