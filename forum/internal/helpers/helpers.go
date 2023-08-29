package helpers

import (
	"encoding/json"
	"strings"

	"github.com/gofrs/uuid"
)

func GenerateUUID() string {
	ID := uuid.Must(uuid.NewV4()).String()
	return ID
}

func DeserializeTags(tagsString string) ([]string, error) {
	// Add surrounding double quotes to create valid JSON
	tagsJSON := `"` + tagsString + `"`

	// Deserialize the JSON string into a slice of strings
	var tags []string
	err := json.Unmarshal([]byte(tagsJSON), &tags)
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func SerializeTags(tags []string) (string, error) {
	// Serialize the slice into a JSON string
	jsonBytes, err := json.Marshal(tags)
	if err != nil {
		return "", err
	}

	// Convert the JSON byte array to a string and remove surrounding double quotes
	tagsString := strings.Trim(string(jsonBytes), `"`)

	return tagsString, nil
}
