package models

type Goal struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Target    string    `json:"target"`
	Progress  float64   `json:"progress"`
	CreatedAt time.Time `json:"created_at"`
}
