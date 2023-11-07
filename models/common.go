package model

import (
	"database/sql"
	"time"
)

type Reponse struct {
	Message    string
	StatusCode int
	Data       any
	HasError   bool
	Errors     []string
}
type DB struct {
	Instance *sql.DB
	Err      error
}
type Category struct {
	Id       any
	Category any
}
type SubmittedData struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type Session struct {
	Token          string
	UserID         string
	ExpirationDate time.Time
	CreationDate   time.Time
}

type UserFeed struct {
	User       User
	Posts      []Post
	ActiveLink string
}
