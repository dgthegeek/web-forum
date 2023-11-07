package helpers

import (
	model "golang-rest-api-starter/models"
	"net/http"
)

func ErrorThrower(err error, message string, statusCode int, w http.ResponseWriter, r *http.Request) {
	response := model.Reponse{
		Message:    message,
		StatusCode: statusCode,
		Data:       nil,
	}
	
	if err != nil {
		ResponseFormatter(
			response, // server response
			"error",  // page where to send response
			w,
			r,
			statusCode,
		)
		return // Stop execution if there's an error
	}
}
