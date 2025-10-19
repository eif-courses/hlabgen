package models

type Review struct {
	ID       int    `json:"id"`
	PaperID  int    `json:"paper_id"`
	Reviewer string `json:"reviewer"`
	Rating   int    `json:"rating"`
	Comments string `json:"comments"`
}
