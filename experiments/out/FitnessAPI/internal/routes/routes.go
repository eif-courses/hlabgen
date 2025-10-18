package routes

import (
	"FitnessAPI/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.POST("/users", handlers.CreateUser)
	r.GET("/users/:id", handlers.GetUser)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	r.POST("/workouts", handlers.CreateWorkout)
	r.GET("/workouts/:id", handlers.GetWorkout)
	r.PUT("/workouts/:id", handlers.UpdateWorkout)
	r.DELETE("/workouts/:id", handlers.DeleteWorkout)

	r.POST("/exercises", handlers.CreateExercise)
	r.GET("/exercises/:id", handlers.GetExercise)
	r.PUT("/exercises/:id", handlers.UpdateExercise)
	r.DELETE("/exercises/:id", handlers.DeleteExercise)

	r.POST("/goals", handlers.CreateGoal)
	r.GET("/goals/:id", handlers.GetGoal)
	r.PUT("/goals/:id", handlers.UpdateGoal)
	r.DELETE("/goals/:id", handlers.DeleteGoal)

	r.POST("/progress", handlers.CreateProgress)
	r.GET("/progress/:id", handlers.GetProgress)
	r.PUT("/progress/:id", handlers.UpdateProgress)
	r.DELETE("/progress/:id", handlers.DeleteProgress)
}
