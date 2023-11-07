package middleware

import (
	"net/http"
)

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log.Println("Request URL:", r.URL.Path)
		next(w, r)
	})
}
