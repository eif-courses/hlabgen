package models

import "time"

type Recipe struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Ingredients []Ingredient `json:"ingredients"`
	Steps       []Step       `json:"steps"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}
type Ingredient struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
}
type Step struct {
	ID          int    `json:"id"`
	RecipeID    int    `json:"recipe_id"`
	Description string `json:"description"`
	Order       int    `json:"order"`
}
