package model

import (
	"time"
)

type User struct {
	ID           string
	FirstName    string
	LastName     string
	Email        string
	Username     string
	Bio          string
	Avatar       string
	Password     string
	CreationDate time.Time
}

