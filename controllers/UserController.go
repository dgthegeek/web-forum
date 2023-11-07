package handlers

import (
	"fmt"
	"golang-rest-api-starter/internals/helpers"
	model "golang-rest-api-starter/models"
	"golang-rest-api-starter/service/crud"
	"log"
	"net/http"
	"strings"
)

type User struct {
	Data      model.UserFeed
	ProfileID string
}

func (u *User) Profile(w http.ResponseWriter, r *http.Request) {
	var (
		p       Post
		comment Comment
	)

	var db, _ = r.Context().Value("db").(model.DB)
	u.ProfileID = strings.TrimPrefix(r.URL.Path, "/users/")
	response := model.Reponse{
		Message:    "OK",
		StatusCode: http.StatusOK,
	} // server response

	// To check what user want to do
	RequestType(db.Instance, r, &response, w)

	columns := []string{}
	conditions := map[string]interface{}{
		"id": u.ProfileID,
	}

	if SelectErr := crud.Select(db.Instance, "User", "", conditions, columns, &u.Data.User.ID, &u.Data.User.FirstName, &u.Data.User.LastName, &u.Data.User.Email, &u.Data.User.Username, &u.Data.User.Bio, &u.Data.User.Avatar, &u.Data.User.Password, &u.Data.User.CreationDate); SelectErr != nil {
		log.Println(SelectErr)
		helpers.ErrorWriter(&response, "Failed to fetch this user. Please try again later", 404)

	}

	u.RequestActivityType(r, db, p, comment)
	response.Data = u.Data

	// When successfully connecting to the database
	helpers.ResponseFormatter(
		response,
		"profile", // page where to send response
		w,
		r, response.StatusCode)
	return
}

func (u *User) RequestActivityType(r *http.Request, db model.DB, p Post, comment Comment) {

	var a = func(postsID []string) {
		var query = "SELECT Post.*, User.username, User.first_name, User.last_name, User.avatar, Category.id AS category_id, Category.category FROM Post JOIN User ON Post.user_id = User.id JOIN PostCategories pc ON Post.id = pc.post_id LEFT JOIN Category ON pc.category_id = Category.id  WHERE Post.id = ? ORDER BY Post.created_at DESC "

		var posts []model.Post
		var singlePost []model.Post
		for _, postid := range FilterUniquePostsID(postsID) {
			singlePost, _ = p.GETPost(r, db.Instance, comment, query, "LIMIT 1", postid)
			posts = append(posts, singlePost...)
		}
		u.Data.Posts = posts
	}

	switch r.URL.Query().Get("activity") {
	case "liked_post":
		u.Data.ActiveLink = "liked_post"

		var postsID []string
		var tableToRetrieveIn = []string{"Postlikes", "Postdislikes", "Commentlikes", "Commentdislikes"}
		for _, table := range tableToRetrieveIn {
			var postid, _ = GETALLPostIDInAnTable(db, u.ProfileID, table)
			postsID = append(postsID, postid...)
		}

		a(postsID)
	case "commented_post":
		u.Data.ActiveLink = "commented_post"
		var postsID, _ = GETALLPostIDInAnTable(db, u.ProfileID, "Comments")
		a(postsID)
	default:
		u.Data.ActiveLink = "my_posts"
		var query = "SELECT Post.*, User.username, User.first_name, User.last_name, User.avatar, Category.id AS category_id, Category.category FROM Post JOIN User ON Post.user_id = User.id JOIN PostCategories pc ON Post.id = pc.post_id LEFT JOIN Category ON pc.category_id = Category.id  WHERE Post.user_id = ? ORDER BY Post.created_at DESC "
		u.Data.Posts, _ = p.GETPost(r, db.Instance, comment, query, "LIMIT 1", u.ProfileID)

	}
}

func GETALLPostIDInAnTable(db model.DB, profileID, table string) ([]string, error) {
	var query = fmt.Sprintf(`SELECT %s.post_id FROM %s WHERE %s.user_id = ?`, table, table, table)
	var rows, err = db.Instance.Query(query, profileID)
	if err != nil {
		return nil, err
	}

	var postsID []string
	for rows.Next() {
		var postID string
		err = rows.Scan(&postID)
		if err != nil {
			return nil, err
		}
		postsID = append(postsID, postID)
	}
	return postsID, nil
}

func FilterUniquePostsID(postsID []string) []string {
	seenPosts := make(map[string]bool)
	uniquePosts := make([]string, 0, len(postsID))

	for _, postID := range postsID {
		if !seenPosts[postID] {
			seenPosts[postID] = true
			uniquePosts = append(uniquePosts, postID)
		}
	}

	return uniquePosts
}
