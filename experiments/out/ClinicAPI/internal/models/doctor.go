package models

import "time"

// Doctor represents a doctor in the clinic.
type Doctor struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Specialty string    `json:"specialty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
