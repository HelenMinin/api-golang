package pkg

import "github.com/google/uuid"

type RequestComment struct {
	Message string `binding:"required"`
	Person  string
}

type Comment struct {
	Id      uuid.UUID
	Message string
	Person  string
}
