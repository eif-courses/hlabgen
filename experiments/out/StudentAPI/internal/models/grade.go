package models

// Grade represents a grade entity.
type Grade struct {
	ID        int     `json:"id"`
	StudentID int     `json:"student_id"`
	CourseID  int     `json:"course_id"`
	Value     float64 `json:"value"`
}
