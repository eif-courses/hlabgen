package models

import "time"

type Patient struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	DOB       time.Time `json:"dob"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
