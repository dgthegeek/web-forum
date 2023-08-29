package middleware

import (
	"net/http"
	"strings"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Cookie"), "=")[0]
		// we can also gram the username and vrify if it is in the database but that will slow down the app
		// this is good for now  in the future we might use stronger libraries
		if authHeader != "sessionID" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
