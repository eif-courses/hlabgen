package models

// Course represents a course entity.
type Course struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Credits int    `json:"credits"`
}
