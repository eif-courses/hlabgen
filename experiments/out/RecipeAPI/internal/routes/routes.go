package routes

import (
	"RecipeAPI/internal/handlers"
	"github.com/gorilla/mux"
)

// Register registers the routes for the application.
func Register(r *mux.Router) {
	r.HandleFunc("/recipes", handlers.CreateRecipe).Methods("POST")
	r.HandleFunc("/recipes", handlers.GetRecipes).Methods("GET")
	r.HandleFunc("/recipes/{id}", handlers.UpdateRecipe).Methods("PUT")
	r.HandleFunc("/recipes/{id}", handlers.DeleteRecipe).Methods("DELETE")
	r.HandleFunc("/ingredients", handlers.CreateIngredient).Methods("POST")
	r.HandleFunc("/ingredients", handlers.GetIngredients).Methods("GET")
	r.HandleFunc("/categories", handlers.CreateCategory).Methods("POST")
	r.HandleFunc("/categories", handlers.GetCategories).Methods("GET")
}
