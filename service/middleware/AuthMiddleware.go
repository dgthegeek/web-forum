package middleware

import (
	"context"
	handlers "golang-rest-api-starter/controllers"
	model "golang-rest-api-starter/models"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session := handlers.Session{}
		// Exclude static files from route verification
		if strings.HasPrefix(r.URL.Path, "/static/") {
			next.ServeHTTP(w, r)
			return
		}
		// Check if we are on authentication routes
		isAuthRoute := strings.HasPrefix(r.URL.Path, "/auth/")

		var db, _ = r.Context().Value("db").(model.DB)
		// check if the session exists in the cookie
		isSessionExist, userID := session.IsSessionExist(&db, r)
		if isAuthRoute && isSessionExist {
			http.Redirect(w, r, "/posts", http.StatusFound)
		} else if !isAuthRoute && isSessionExist {
			r = r.WithContext(context.WithValue(r.Context(), "isAuthentificated", true))
			// I wanna know inside my back end wether the actual has like or not an post.

			r = r.WithContext(context.WithValue(r.Context(), "userID", userID))
			next.ServeHTTP(w, r)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), "isAuthentificated", false))
		next.ServeHTTP(w, r)
	}
}
