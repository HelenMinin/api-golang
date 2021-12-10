package db

import (
	"facec/blog/pkg"
)

var commentPostMap = make(map[string]map[string]*pkg.Comment)

type CommentRepository struct {
}

func (CommentRepository) InsertComment(postId string, comment *pkg.Comment) {
	commentMap := commentPostMap[postId]
	commentMap[comment.Id.String()] = comment
}

func (CommentRepository) GetComments(postId string) []*pkg.Comment {
	comments := make([]*pkg.Comment, 0)
	for _, v := range commentPostMap[postId] {
		comments = append(comments, v)
	}

	return comments
}
