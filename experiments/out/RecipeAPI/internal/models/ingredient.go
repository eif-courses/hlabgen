package models

// Ingredient represents an ingredient entity.
type Ingredient struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
}
