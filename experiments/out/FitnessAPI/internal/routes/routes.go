package routes

import (
	"FitnessAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/workouts", handlers.CreateWorkout).Methods("POST")
	r.HandleFunc("/exercises", handlers.CreateExercise).Methods("POST")
	r.HandleFunc("/goals", handlers.CreateGoal).Methods("POST")
	r.HandleFunc("/progress", handlers.CreateProgress).Methods("POST")
}
