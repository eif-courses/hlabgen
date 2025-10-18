package models

import "time"

type Loan struct {
	ID         int        `json:"id"`
	BookID     int        `json:"book_id"`
	UserID     int        `json:"user_id"`
	LoanedAt   time.Time  `json:"loaned_at"`
	ReturnedAt *time.Time `json:"returned_at,omitempty"`
}
