package routes

import (
	"SurveyAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	// Survey routes
	r.HandleFunc("/surveys", handlers.CreateSurvey).Methods("POST")
	r.HandleFunc("/surveys", handlers.GetSurveys).Methods("GET")
	r.HandleFunc("/surveys/{id}", handlers.GetSurvey).Methods("GET")
	r.HandleFunc("/surveys/{id}", handlers.UpdateSurvey).Methods("PUT")
	r.HandleFunc("/surveys/{id}", handlers.DeleteSurvey).Methods("DELETE")
}
