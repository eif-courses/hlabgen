package models

type Grade struct {
	ID        int     `json:"id"`
	StudentID int     `json:"student_id"`
	CourseID  int     `json:"course_id"`
	Score     float64 `json:"score"`
}
