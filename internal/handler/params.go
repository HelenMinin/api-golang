package handler

import (
	"facec/blog/internal/container"
	"facec/blog/pkg"
	"fmt"

	"github.com/gin-gonic/gin"
)

func findPost(c *gin.Context, paramName string) (*pkg.Post, *pkg.ResponseError) {
	id := c.Param(paramName)

	post := container.PostRepository.FindById(id)
	if post == nil {
		return nil, &pkg.ResponseError{
			Cause:   paramName + " not found",
			Message: fmt.Sprintf("id %s not found", id),
		}
	}

	return post, nil
}
