package models

// Project represents a project entity.
type Project struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Task represents a task entity.
type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	ProjectID int    `json:"project_id"`
	Status    string `json:"status"`
}

// Milestone represents a milestone entity.
type Milestone struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ProjectID int    `json:"project_id"`
	DueDate   string `json:"due_date"`
}
