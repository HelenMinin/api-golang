package pkg

import "time"

type RequestPost struct {
	Title string
	Body  string
	User  string
}

type Post struct {
	Title    string
	Body     string
	User     string
	DateTime time.Time
}
