package handler

import (
	"api/golang/pkg"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

var posts = make([]pkg.Post, 3)

func GetPosts(c *gin.Context) {
	c.JSON(200, posts)
}

func CreatePost(c *gin.Context) {
	var requestPost pkg.RequestPost

	if err := c.ShouldBindJSON(&requestPost); err != nil {
		log.Println("[WARN] - invalid json ", err)
		c.JSON(400, pkg.NewResponseError("Invalid JSON", err))

		return
	}

	post := pkg.Post{
		Title:    requestPost.Title,
		Body:     requestPost.Body,
		User:     requestPost.User,
		DateTime: time.Now(),
	}

	posts = append(posts, post)

	log.Println(fmt.Sprintln("post &s created", requestPost))

	c.JSON(201, posts)
}
