package routes

import (
	"TaskAPI/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.POST("/tasks", handlers.CreateTask)
	r.GET("/tasks", handlers.GetTasks)
	r.PUT("/tasks/:id", handlers.UpdateTask)
	r.DELETE("/tasks/:id", handlers.DeleteTask)

	r.POST("/projects", handlers.CreateProject)
	r.GET("/projects", handlers.GetProjects)
	r.PUT("/projects/:id", handlers.UpdateProject)
	r.DELETE("/projects/:id", handlers.DeleteProject)

	r.POST("/users", handlers.CreateUser)
	r.GET("/users", handlers.GetUsers)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	r.POST("/teams", handlers.CreateTeam)
	r.GET("/teams", handlers.GetTeams)
	r.PUT("/teams/:id", handlers.UpdateTeam)
	r.DELETE("/teams/:id", handlers.DeleteTeam)
}
