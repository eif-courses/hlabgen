package models

import "encoding/json"

// Recipe represents a recipe in the system.
type Recipe struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	CategoryID  int          `json:"category_id"`
	Ingredients []Ingredient `json:"ingredients"`
}

// Ingredient represents an ingredient in a recipe.
type Ingredient struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
}

// Category represents a category of recipes.
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ToJSON converts a Recipe to JSON.
func (r *Recipe) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

// FromJSON converts JSON to a Recipe.
func FromJSON(data []byte) (*Recipe, error) {
	var r Recipe
	err := json.Unmarshal(data, &r)
	return &r, err
}
