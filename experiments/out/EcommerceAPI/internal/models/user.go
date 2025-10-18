package models

import "encoding/json"

// User represents a user in the e-commerce system.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// ToJSON converts a User to JSON.
func (u *User) ToJSON() ([]byte, error) {
	return json.Marshal(u)
}
