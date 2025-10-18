package routes

import (
	"RecipeAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register() {
	r.HandleFunc("/recipes", handlers.CreateRecipe).Methods("POST")
	r.HandleFunc("/recipes", handlers.GetRecipes).Methods("GET")
	r.HandleFunc("/recipes/{id}", handlers.UpdateRecipe).Methods("PUT")
	r.HandleFunc("/recipes/{id}", handlers.DeleteRecipe).Methods("DELETE")
}
