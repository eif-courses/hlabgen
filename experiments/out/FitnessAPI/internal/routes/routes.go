package routes

import (
	"FitnessAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register() {
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/workouts", handlers.CreateWorkout).Methods("POST")
	r.HandleFunc("/workouts/{id:[0-9]+}", handlers.GetWorkout).Methods("GET")
	r.HandleFunc("/workouts/{id:[0-9]+}", handlers.UpdateWorkout).Methods("PUT")
	r.HandleFunc("/workouts/{id:[0-9]+}", handlers.DeleteWorkout).Methods("DELETE")
	r.HandleFunc("/exercises", handlers.CreateExercise).Methods("POST")
	r.HandleFunc("/exercises/{id:[0-9]+}", handlers.GetExercise).Methods("GET")
	r.HandleFunc("/exercises/{id:[0-9]+}", handlers.UpdateExercise).Methods("PUT")
	r.HandleFunc("/exercises/{id:[0-9]+}", handlers.DeleteExercise).Methods("DELETE")
	r.HandleFunc("/goals", handlers.CreateGoal).Methods("POST")
	r.HandleFunc("/goals/{id:[0-9]+}", handlers.GetGoal).Methods("GET")
	r.HandleFunc("/goals/{id:[0-9]+}", handlers.UpdateGoal).Methods("PUT")
	r.HandleFunc("/goals/{id:[0-9]+}", handlers.DeleteGoal).Methods("DELETE")
}
