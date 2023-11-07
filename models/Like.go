package model

import "time"

type Like struct {
	ID           string
	AuthorID     string
	PostID       string
	CreationDate time.Time
	UserName 	 string
	FirstName  string // Field to store the username
	LastName   string // Field to store the username
	Avatar     string // Field to store the username
}
