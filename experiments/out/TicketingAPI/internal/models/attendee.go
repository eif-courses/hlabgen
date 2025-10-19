package models

type Attendee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	TicketID int    `json:"ticket_id"`
}
