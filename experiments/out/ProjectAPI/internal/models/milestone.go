package models

type Milestone struct {
	ID        int    `json:"id"`
	ProjectID int    `json:"project_id"`
	Title     string `json:"title"`
	DueDate   string `json:"due_date"`
}
