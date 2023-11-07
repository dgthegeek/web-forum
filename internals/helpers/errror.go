package helpers

import (
	model "golang-rest-api-starter/models"
)

func ErrorWriter(response *model.Reponse, message string, statusCode int) {
	response.Errors = append(response.Errors, message) // the error message to send the client
	response.StatusCode = statusCode                   // error that you want to assign to the response header
	response.HasError = true                           // the error modal on the UI depends on this value
}
