package models

type Task struct {
	ID        int    `json:"id"`
	ProjectID int    `json:"project_id"`
	Title     string `json:"title"`
	Status    string `json:"status"`
}
