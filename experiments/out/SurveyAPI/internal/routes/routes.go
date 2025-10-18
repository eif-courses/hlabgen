package routes

import (
	"SurveyAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register() {
	router.HandleFunc("/surveys", handlers.CreateSurvey).Methods("POST")
	router.HandleFunc("/surveys", handlers.GetSurveys).Methods("GET")
	router.HandleFunc("/surveys/{id}", handlers.UpdateSurvey).Methods("PUT")
	router.HandleFunc("/surveys/{id}", handlers.DeleteSurvey).Methods("DELETE")
	router.HandleFunc("/questions", handlers.CreateQuestion).Methods("POST")
	router.HandleFunc("/questions", handlers.GetQuestions).Methods("GET")
	router.HandleFunc("/questions/{id}", handlers.UpdateQuestion).Methods("PUT")
	router.HandleFunc("/questions/{id}", handlers.DeleteQuestion).Methods("DELETE")
	router.HandleFunc("/responses", handlers.CreateResponse).Methods("POST")
	router.HandleFunc("/responses", handlers.GetResponses).Methods("GET")
	router.HandleFunc("/responses/{id}", handlers.UpdateResponse).Methods("PUT")
	router.HandleFunc("/responses/{id}", handlers.DeleteResponse).Methods("DELETE")
}
