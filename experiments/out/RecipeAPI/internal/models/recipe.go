package models

import "time"

// Recipe represents a recipe with its details.
type Recipe struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Instructions string    `json:"instructions"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Ingredient represents an ingredient used in a recipe.
type Ingredient struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
	RecipeID int    `json:"recipe_id"`
}

// Category represents a category for recipes.
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
