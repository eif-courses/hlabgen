package routes

import (
	"FitnessAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	// User routes
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	// Workout routes
	r.HandleFunc("/workouts", handlers.CreateWorkout).Methods("POST")
	// Exercise routes
	r.HandleFunc("/exercises", handlers.CreateExercise).Methods("POST")
	// Goal routes
	r.HandleFunc("/goals", handlers.CreateGoal).Methods("POST")
}
