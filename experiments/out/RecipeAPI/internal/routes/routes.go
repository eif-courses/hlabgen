package routes

import (
	"RecipeAPI/internal/handlers"
	"github.com/gorilla/mux"
)

// Register sets up the routes for the application.
func Register(r *mux.Router) {
	// Recipe routes
	r.HandleFunc("/recipes", handlers.ListRecipes).Methods("GET")
	r.HandleFunc("/recipes", handlers.CreateRecipe).Methods("POST")
	r.HandleFunc("/recipes/{id}", handlers.GetRecipe).Methods("GET")
	r.HandleFunc("/recipes/{id}", handlers.UpdateRecipe).Methods("PUT")
	r.HandleFunc("/recipes/{id}", handlers.DeleteRecipe).Methods("DELETE")

	// Ingredient routes
	r.HandleFunc("/ingredients", handlers.ListIngredients).Methods("GET")
	r.HandleFunc("/ingredients", handlers.CreateIngredient).Methods("POST")
	r.HandleFunc("/ingredients/{id}", handlers.GetIngredient).Methods("GET")
	r.HandleFunc("/ingredients/{id}", handlers.UpdateIngredient).Methods("PUT")
	r.HandleFunc("/ingredients/{id}", handlers.DeleteIngredient).Methods("DELETE")

	// Category routes
	r.HandleFunc("/categories", handlers.ListCategories).Methods("GET")
	r.HandleFunc("/categories", handlers.CreateCategory).Methods("POST")
	r.HandleFunc("/categories/{id}", handlers.GetCategory).Methods("GET")
	r.HandleFunc("/categories/{id}", handlers.UpdateCategory).Methods("PUT")
	r.HandleFunc("/categories/{id}", handlers.DeleteCategory).Methods("DELETE")
}
