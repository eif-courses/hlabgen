package models

// Step represents a step in a recipe.
type Step struct {
	ID          int    `json:"id"`
	Order       int    `json:"order"`
	Instruction string `json:"instruction"`
}
