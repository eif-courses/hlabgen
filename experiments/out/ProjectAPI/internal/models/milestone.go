package models

type Milestone struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	DueDate   string `json:"due_date"`
	ProjectID int    `json:"project_id"`
}
