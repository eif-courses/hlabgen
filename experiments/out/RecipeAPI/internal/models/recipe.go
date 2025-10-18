package models


// Recipe represents a recipe entity.
type Recipe struct {
ID          int         `json:"id"`
    Title       string      `json:"title"`
    Description string      `json:"description"`
    Ingredients [Ingredient `json:"ingredients"`
    Steps       []Step      `json:"steps"`
}
// Ingredient represents an ingredient entity.
type Ingredient struct {
ID       int    `json:"id"`
    Name     string `json:"name"`
    Quantity string `json:"quantity"`
}
// Step represents a step in a recipe.
type Step struct {
ID      int    `json:"id"`
    Order   int    `json:"order"`
    Instruction string `json:"instruction"`
}