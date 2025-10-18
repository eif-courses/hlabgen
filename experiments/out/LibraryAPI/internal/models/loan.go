package models

import "time"

// Loan represents a loan of a book by a user.
type Loan struct {
	ID       int       `json:"id"`
	BookID   int       `json:"book_id"`
	UserID   int       `json:"user_id"`
	LoanedAt time.Time `json:"loaned_at"`
	DueDate  time.Time `json:"due_date"`
}
