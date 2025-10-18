package models

// Ingredient represents an ingredient used in a recipe.
type Ingredient struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
	RecipeID int    `json:"recipe_id"`
}
