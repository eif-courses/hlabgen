package models

// Review represents a review entity.
type Review struct {
	ID      int    `json:"id"`
	MovieID int    `json:"movie_id"`
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}
