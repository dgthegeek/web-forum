package middleware

import (
	"context"
	db "golang-rest-api-starter/internals/config/database"
	"golang-rest-api-starter/internals/helpers"
	model "golang-rest-api-starter/models"
	"net/http"
)

/*
This middleware has as purpose to send an instance of
the database to the routes
*/
func DBMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		DB := model.DB{Instance: nil, Err: nil}

		database := &db.Config{
			Driver: "sqlite3",
			Name:   "forum.db",
		}
		/* Check if a database connection is already opened
		to avoid opening multiple databases connections */
		if DB.Instance == nil {
			DB.Instance, DB.Err = database.Init()
			// Throw the error page when we have a database issue
			if DB.Err != nil {
				helpers.ErrorThrower(DB.Err, "Something went wrong if the problem persists, please contact the us", 500, w, r)
				return
			}
		}
		// Create context to pass the database instance across all routes
		var ctx = context.Background()
		ctx = context.WithValue(r.Context(), "db", DB)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
