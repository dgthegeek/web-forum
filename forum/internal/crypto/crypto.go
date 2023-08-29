package crypto

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateRandomSessionID generates a cryptographically secure random session ID.
func GenerateRandomSessionID(length int) (string, error) {
	// Determine the number of bytes needed to represent the session ID
	bytesNeeded := (length + 3) / 4 * 3

	// Generate random bytes
	randomBytes := make([]byte, bytesNeeded)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}

	// Encode the random bytes to a URL-safe base64 string
	sessionID := base64.URLEncoding.EncodeToString(randomBytes)[:length]

	return sessionID, nil
}
