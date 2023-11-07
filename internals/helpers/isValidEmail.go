package helpers

import "regexp"

// The purpose of this function is to check if the provided email is valid via regexp
func IsValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression
	re := regexp.MustCompile(emailRegex)

	// Test the email against the regex
	return re.MatchString(email)
}
