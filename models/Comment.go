package model

import "time"

type Comment struct {
	ID           string
	Content      string
	AuthorID     string
	PostID       string
	UserName 	 string
	CreationDate time.Time
	FirstName  string // Field to store the username
	LastName   string // Field to store the username
	Avatar     string // Field to store the username
	LikesCount, DisLikesCount int
	
}
