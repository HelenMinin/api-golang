package http

import (
	"api/golang/internal/handler"

	"github.com/gin-gonic/gin"
)

func StartServer() error {
	engine := gin.Default()

	engine.GET("/health", handler.HealthCheck)
	engine.GET("/posts", handler.GetPosts)
	engine.POST("/posts", handler.CreatePost)

	return engine.Run(":3000")
}
