package router

import (
	"errors"

	handlers "golang-rest-api-starter/controllers"
	"golang-rest-api-starter/internals/helpers"
	"net/http"
	"strings"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc
type NewRouter struct {
	Middlewares []Middleware
}

func (router *NewRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// before going in any endpoints make sure the method is "POST" or "GET".
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		helpers.ErrorThrower(errors.New("Error method"), "Method not allowed.", http.StatusMethodNotAllowed, w, r)
		return
	}

	post := handlers.Post{}
	auth := handlers.Auth{}
	user := handlers.User{}

	handler := func(w http.ResponseWriter, r *http.Request) {
		switch {
		// Home Page
		case r.URL.Path == "/":
			handlers.HomeHandler(w, r)

		// Posts routes
		case r.URL.Path == "/posts":
			post.Feed(w, r)
		case strings.HasPrefix(r.URL.Path, "/posts/"):
			post.Single(w, r)

		// Auth routes
		case r.URL.Path == "/auth/login":
			auth.Login(w, r)
		case r.URL.Path == "/auth/register":
			auth.Register(w, r)
		// User routes
		case strings.HasPrefix(r.URL.Path, "/users/"):
			user.Profile(w, r)
		// Static files
		case strings.HasPrefix(r.URL.Path, "/static/"):
			fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
			fs.ServeHTTP(w, r)

		// When the URL doesn't exist
		default:
			func(w http.ResponseWriter, _ *http.Request) {
				helpers.ErrorThrower(errors.New("Not Found"), "Oops! This page doesn't that you looking for does not exist. It might be move or delete", 404, w, r)
			}(w, r)
		}
	}

	for _, mw := range router.Middlewares {
		handler = mw(handler)
	}
	handler(w, r)

}
