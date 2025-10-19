package models

type Question struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
	Type     string `json:"type"`
	SurveyID int    `json:"survey_id"`
}
