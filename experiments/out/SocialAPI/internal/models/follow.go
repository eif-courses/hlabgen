package models

// Follow represents a follow relationship between users.
type Follow struct {
	ID       int `json:"id"`
	Follower int `json:"follower_id"`
	Followed int `json:"followed_id"`
}
