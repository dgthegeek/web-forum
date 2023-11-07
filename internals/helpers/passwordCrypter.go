package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

// The purpose of this function is to encrypt password before store it in the database
func PasswordEncrypter(password string) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil
	}
	return hashedPassword
}

func PasswordDecrypter(userPassword, submittedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(submittedPassword)) == nil
}
