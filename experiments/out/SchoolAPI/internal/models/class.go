package models

import "time"

type Class struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	TeacherID int       `json:"teacher_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
