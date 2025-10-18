package routes

import (
    "github.com/gin-gonic/gin"
    "BlogAPI/internal/handlers"
)

func Register(router *gin.Engine) {
    posts := router.Group("/posts") {
        router.POST("/", handlers.CreatePost)
        router.GET("/", handlers.GetPosts)
        router.GET("/:id", handlers.GetPost)
        router.PUT("/:id", handlers.UpdatePost)
        router.DELETE("/:id", handlers.DeletePost)
    }
    comments := router.Group("/comments") {
        router.POST("/", handlers.CreateComment)
        router.GET("/", handlers.GetComments)
        router.GET("/:id", handlers.GetComment)
        router.PUT("/:id", handlers.UpdateComment)
        router.DELETE("/:id", handlers.DeleteComment)
    }
    authors := router.Group("/authors") {
        router.POST("/", handlers.CreateAuthor)
        router.GET("/", handlers.GetAuthors)
        router.GET("/:id", handlers.GetAuthor)
        router.PUT("/:id", handlers.UpdateAuthor)
        router.DELETE("/:id", handlers.DeleteAuthor)
    }
}
