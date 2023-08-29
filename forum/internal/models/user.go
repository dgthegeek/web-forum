// models/user.go

package models

type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	// Add more fields as needed
}

type ResetRequest struct {
	Email string `json:"email"`
	NewPassword string `json:"newpassword"`

}