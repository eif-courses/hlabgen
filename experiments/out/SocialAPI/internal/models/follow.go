package models

type Follow struct {
	ID       int `json:"id"`
	Follower int `json:"follower_id"`
	Followed int `json:"followed_id"`
}
