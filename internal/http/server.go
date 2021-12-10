package http

import (
	"facec/blog/internal/handler"

	"github.com/gin-gonic/gin"
)

func StartServer() error {
	engine := gin.Default()
	gin.EnableJsonDecoderDisallowUnknownFields()

	engine.GET("/health", handler.HealthCheck)

	posts := engine.Group("/posts")
	posts.GET("/", handler.GetPosts)
	posts.POST("/", handler.CreatePost)

	posts.GET("/:id", handler.GetPost)
	posts.PUT("/:id", handler.UpdatePost)
	posts.PATCH("/:id", handler.PartialUpdatePost)
	posts.DELETE("/:id", handler.DeletePost)

	posts.POST("/:postId/comments", handler.CreateComments)
	//posts.GET("/:id/comments", handler.GetComments)

	//posts.GET("/:id/comments/:id", handler.CreateComment)

	return engine.Run()
}
