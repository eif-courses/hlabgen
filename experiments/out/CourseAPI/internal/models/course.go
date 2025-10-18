package models

// Course represents a course entity.
type Course struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Lessons     []Lesson `json:"lessons"`
}

// Lesson represents a lesson entity.
type Lesson struct {
	ID       int    `json:"id"`
	CourseID int    `json:"course_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

// Enrollment represents an enrollment entity.
type Enrollment struct {
	ID       int `json:"id"`
	CourseID int `json:"course_id"`
	UserID   int `json:"user_id"`
}
