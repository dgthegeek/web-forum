package handlers

import (
	"database/sql"
	"golang-rest-api-starter/internals/helpers"
	model "golang-rest-api-starter/models"
	"net/http"
)

// Homepage handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	helpers.ResponseFormatter(model.Reponse{StatusCode: http.StatusOK}, "home", w, r, http.StatusOK)
}

func RequestType(db *sql.DB, r *http.Request, response *model.Reponse, w http.ResponseWriter) {
	var (
		comment = Comment{}
		action  = Action{}
		post    = Post{}
		session = Session{}
	)

	if r.Method == http.MethodPost {
		if err := r.ParseMultipartForm(20 * 1024 * 1024); err != nil {
			helpers.ErrorWriter(response, "Cannot perform this operation. Please try later.", 500)
		}
	
		switch r.FormValue("post-type") {
		case "post":
			post.Create(r, db, response, w)
		case "comment":
			comment.Create(r, db, response)
		case "like":
			action.Create(r, db, response, "Likes", "disLikes")
		case "dislike":
			action.Create(r, db, response, "DisLikes", "likes")
		case "log-out":
			if err := session.LogOut(r); err != nil {
				helpers.ErrorWriter(response, "Something went wrong. Please try later.", 500)
			}
		}
	}
}
