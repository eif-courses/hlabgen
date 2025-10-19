package routes

import (
	"RecipeAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	// Recipe routes
	r.HandleFunc("/recipes", handlers.CreateRecipe).Methods("POST")
	r.HandleFunc("/recipes", handlers.GetRecipes).Methods("GET")
	r.HandleFunc("/recipes/{id}", handlers.GetRecipe).Methods("GET")
	r.HandleFunc("/recipes/{id}", handlers.UpdateRecipe).Methods("PUT")
	r.HandleFunc("/recipes/{id}", handlers.DeleteRecipe).Methods("DELETE")

	// Ingredient routes
	r.HandleFunc("/ingredients", handlers.CreateIngredient).Methods("POST")
	r.HandleFunc("/ingredients", handlers.GetIngredients).Methods("GET")
	r.HandleFunc("/ingredients/{id}", handlers.GetIngredient).Methods("GET")
	r.HandleFunc("/ingredients/{id}", handlers.UpdateIngredient).Methods("PUT")
	r.HandleFunc("/ingredients/{id}", handlers.DeleteIngredient).Methods("DELETE")

	// Step routes
	r.HandleFunc("/steps", handlers.CreateStep).Methods("POST")
	r.HandleFunc("/steps", handlers.GetSteps).Methods("GET")
	r.HandleFunc("/steps/{id}", handlers.GetStep).Methods("GET")
	r.HandleFunc("/steps/{id}", handlers.UpdateStep).Methods("PUT")
	r.HandleFunc("/steps/{id}", handlers.DeleteStep).Methods("DELETE")
}
