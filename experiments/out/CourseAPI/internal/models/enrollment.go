package models

// Enrollment represents a user's enrollment in a course.
type Enrollment struct {
	ID       int `json:"id"`
	CourseID int `json:"course_id"`
	UserID   int `json:"user_id"`
}
