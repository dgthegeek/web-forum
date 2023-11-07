package helpers

import (
	"errors"
	"fmt"
	model "golang-rest-api-starter/models"
	"strings"
)

// The purpose of this function is to validate datas passed as parameters and return a response
func ValidateForm(submittedData map[string]interface{}, response *model.Reponse) error {
	if len(submittedData) == 0 {
		ErrorWriter(response, "Certain fields cannot be empty", 400)
	}
	for key, value := range submittedData {
		if key == "password" {
			haveSpace := strings.Contains(value.(string), " ")
			if haveSpace {
				ErrorWriter(response, "Password cannot contain white space", 400)
			}
		}
		// Unrequired fields are ignored during form validation
		if value == "" && key != "image" && key != "bio" {
			// Unfilled fields ❌
			message := fmt.Sprintf("*the %s field cannot be empty.", key)
			ErrorWriter(response, message, 400)
		} else if key == "email" && !IsValidEmail(value.(string)) {
			// Invalid email format ❌
			message := "*Please enter a valid email address"
			ErrorWriter(response, message, 400)
		}
	}

	return nil
}

func IsRequiredFeildsExits(submittedData map[string]interface{}, fields ...string) error {
	for _, field := range fields {
		var _, ok = submittedData[field]
		if !ok {
			return errors.New("Some Required Fields are not found.")
		}
	}
	return nil
}
