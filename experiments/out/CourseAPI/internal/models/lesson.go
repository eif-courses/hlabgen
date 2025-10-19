package models

type Lesson struct {
	ID       int    `json:"id"`
	CourseID int    `json:"course_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}
