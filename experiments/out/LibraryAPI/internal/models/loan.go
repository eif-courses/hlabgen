package models

type Loan struct {
	ID       int    `json:"id"`
	BookID   int    `json:"book_id"`
	MemberID int    `json:"member_id"`
	DueDate  string `json:"due_date"`
}
