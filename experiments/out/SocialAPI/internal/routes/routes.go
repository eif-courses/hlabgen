package routes

import (
	"SocialAPI/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.POST("/users", handlers.CreateUser)
	r.GET("/users/:id", handlers.GetUser)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	r.POST("/posts", handlers.CreatePost)
	r.GET("/posts/:id", handlers.GetPost)
	r.PUT("/posts/:id", handlers.UpdatePost)
	r.DELETE("/posts/:id", handlers.DeletePost)

	r.POST("/comments", handlers.CreateComment)
	r.GET("/comments/:id", handlers.GetComment)
	r.PUT("/comments/:id", handlers.UpdateComment)
	r.DELETE("/comments/:id", handlers.DeleteComment)

	r.POST("/likes", handlers.LikePost)
	r.DELETE("/likes", handlers.UnlikePost)

	r.POST("/follows", handlers.FollowUser)
	r.DELETE("/follows", handlers.UnfollowUser)

	r.GET("/notifications", handlers.GetNotifications)
}
