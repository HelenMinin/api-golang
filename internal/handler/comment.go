package handler

import (
	"facec/blog/internal/container"
	"facec/blog/pkg"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const paramPostId = "postId"

func CreateComment(c *gin.Context) {
	post, responseError := findPost(c, paramPostId)

	if responseError != nil {
		c.JSON(404, responseError)
		return
	}
	var requestComment *pkg.Comment
	if err := c.ShouldBindJSON(&requestComment); err != nil {
		c.JSON(404, &pkg.ResponseError{
			Cause:   "Invalid request body",
			Message: err.Error(),
		})

		return
	}

	comment := pkg.Comment{
		Id:      uuid.New(),
		Message: requestComment.Message,
		Person:  requestComment.Person,
	}

	container.CommentRepository.InsertComment(post.Id.String(), &comment)
	c.JSON(201, comment)
}

func GetCommentsByPost(c *gin.Context) {
	post, responseError := findPost(c, paramPostId)

	if responseError != nil {
		c.JSON(404, responseError)
		return
	}

	comments := container.CommentRepository.GetComments(post.Id.String())

	c.JSON(200, comments)
}
