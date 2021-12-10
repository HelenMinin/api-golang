package pkg

import (
	"time"

	"github.com/google/uuid"
)

type RequestPost struct {
	Title string `binding:"required"`
	Body  string `binding:"required"`
	User  string `binding:"required"`
}

type PartialRequestPost struct {
	Title *string
	Body  *string
	User  *string
}

type Post struct {
	Id       uuid.UUID
	Title    string
	Body     string
	User     string
	DateTime time.Time
}
