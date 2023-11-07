package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	internals "golang-rest-api-starter/internals/config/database"
	"golang-rest-api-starter/internals/helpers"
	model "golang-rest-api-starter/models"
	"golang-rest-api-starter/service/crud"
	"log"
	"net/http"
	"strings"
)

type Post struct{}

// Method that retreives all posts
func (p *Post) Feed(w http.ResponseWriter, r *http.Request) {
	var db, _ = r.Context().Value("db").(model.DB)

	// This is the default response sent to the client
	response := model.Reponse{
		Message:    "OK",
		StatusCode: http.StatusOK,
	}
	internals.TablesCreation(db.Instance)

	// To check what user want to do
	RequestType(db.Instance, r, &response, w)

	topic := r.URL.Query().Get("topic")
	comment := Comment{}
	// The sql query depends on
	query := FilterPost(topic)

	posts, fetchPostErr := p.GETPost(r, db.Instance, comment, query, "LIMIT 1", "")
	if fetchPostErr != nil {
		log.Println(fetchPostErr)
		helpers.ErrorWriter(&response, "Failed to fetch your feed", 500)
	}

	response.Data = posts
	helpers.ResponseFormatter(
		response,
		"posts",
		w, r, response.StatusCode)
}

// Logic to get a single Post
func (p *Post) Single(w http.ResponseWriter, r *http.Request) {
	comment := Comment{}
	postID := strings.TrimPrefix(r.URL.Path, "/posts/")
	var db, _ = r.Context().Value("db").(model.DB)
	internals.TablesCreation(db.Instance)

	var response = model.Reponse{
		Message:    "OK",
		StatusCode: http.StatusOK,
		HasError:   false,
	}

	// To check what user want to do
	RequestType(db.Instance, r, &response, w)

	var query = "SELECT Post.*, User.username, User.first_name, User.last_name, User.avatar, Category.id AS category_id, Category.category FROM Post JOIN User ON Post.user_id = User.id JOIN PostCategories pc ON Post.id = pc.post_id LEFT JOIN Category ON pc.category_id = Category.id  WHERE Post.id = ? ORDER BY Post.created_at DESC "
	posts, Err := p.GETPost(r, db.Instance, comment, query, "", postID)
	if Err != nil {
		helpers.ErrorWriter(&response, "Cannot fetch this post.", http.StatusInternalServerError)
	}
	if len(posts) == 0 {
		response.StatusCode = http.StatusNotFound
	}
	if response.StatusCode == http.StatusOK && len(posts) > 0 {
		response.Data = posts[0]
	}

	// When successfully connecting to the database
	helpers.ResponseFormatter(response, "post", w, r, response.StatusCode)
}

// Logic for Post creation
func (p *Post) Create(r *http.Request, db *sql.DB, response *model.Reponse, w http.ResponseWriter) {

	datas := helpers.GetFormData(r)
	var userID, _ = r.Context().Value("userID").(string)
	if userID == "" {
		helpers.ErrorWriter(response, "You don't have authorization to perfom this action", 403)
	}
	helpers.ValidateForm(datas, response)
	var err = helpers.IsRequiredFeildsExits(datas, "title", "description", "category")
	if err != nil {
		helpers.ErrorWriter(response, err.Error(), http.StatusBadRequest)
	}
	if response.StatusCode == http.StatusOK {
		uploadedImg, Err := helpers.ImageUploader(w, r)

		if Err != nil && Err.Error() != "http: no such file" {
			helpers.ErrorWriter(response, Err.Error(), 400)
		}

		if len(datas["description"].(string)) > 255 {
			helpers.ErrorWriter(response, "Your post description cannot exceed 255 characters.", http.StatusBadRequest)
		}
		if len(datas["title"].(string)) > 100 {
			helpers.ErrorWriter(response, "Your post title cannot exceed 100 characters.", http.StatusBadRequest)
		}

		if response.StatusCode == http.StatusOK {
			columnsToInsert := []string{"image", "title", "content", "user_id"}
			publishedPostID, InsertionErr := crud.Insert(db, "Post", columnsToInsert, uploadedImg, datas["title"], datas["description"], userID)
			if InsertionErr != nil {
				log.Println(InsertionErr)
				helpers.ErrorWriter(response, "Failed to publish your post. Try again later", http.StatusBadRequest)
			}
			if response.StatusCode == http.StatusOK {
				// Insert selected categories into the database
				// Assuming datas is a map[string]interface{}
				categoryData := datas["category"]

				switch categoryData := categoryData.(type) {
				case []string:
					// It's a slice of strings, loop through and process each string.
					for _, v := range categoryData {
						columns := []string{"post_id", "category_id"}
						_, Err := crud.Insert(db, "PostCategories", columns, publishedPostID, v)
						if Err != nil {
							helpers.ErrorWriter(response, "*Failed while creating category", http.StatusBadRequest)
						}
					}
				case string:
					// It's a single string, you can process it as needed.
					columns := []string{"post_id", "category_id"}
					_, Err := crud.Insert(db, "PostCategories", columns, publishedPostID, categoryData)
					if Err != nil {
						helpers.ErrorWriter(response, "*Failed while creating category", http.StatusBadRequest)
					}
				}
			}
		}
	}
}

func (p *Post) GETPost(r *http.Request, db *sql.DB, comment Comment, query, limit, condition string) ([]model.Post, error) {
	rows, err := db.Query(query, condition)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []model.Post
	for rows.Next() {

		var post model.Post
		var category model.Category

		if err := rows.Scan(
			&post.ID, &post.Image, &post.Title, &post.Content, &post.UserID, &post.LikesCount, &post.DisLikesCount, &post.CommentsCounts, &post.CreatedAt, &post.Username, &post.FirstName, &post.LastName, &post.Avatar, &category.Id, &category.Category); err != nil {
			return nil, err
		}

		// get the comments related to the post
		post.Comments, err = comment.GetPostComments(db, post.ID, limit)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		if existingPost, exists := findPostByID(posts, post.ID); !exists {
			post = model.Post{ID: post.ID, Image: post.Image, Title: post.Title, Content: post.Content, UserID: post.UserID, LikesCount: post.LikesCount, DisLikesCount: post.DisLikesCount, CommentsCounts: post.CommentsCounts, Comments: post.Comments, CreatedAt: post.CreatedAt, Username: post.Username, FirstName: post.FirstName, LastName: post.LastName, Avatar: post.Avatar,
				Categories: []model.Category{category},
			}
			posts = append(posts, post)
		} else {
			existingPost.Categories = append(existingPost.Categories, category)
		}
	}
	return posts, nil
}

func findPostByID(posts []model.Post, id string) (*model.Post, bool) {
	for i, post := range posts {
		if post.ID == id {
			return &posts[i], true
		}
	}
	return nil, false
}

func FilterPost(topic string) string {
	query := "SELECT Post.*, User.username, User.first_name, User.last_name, User.avatar, Category.id AS category_id, Category.category FROM Post JOIN User ON Post.user_id = User.id JOIN PostCategories pc ON Post.id = pc.post_id LEFT JOIN Category ON pc.category_id = Category.id"
	if topic != "" {
		query = fmt.Sprintf("%s WHERE Category.category ='%s' ORDER BY Post.created_at DESC", query, topic)
	} else {
		query = fmt.Sprintf("%s  ORDER BY Post.created_at DESC", query)
	}
	return query
}
func TrendingFilter(filter string) string {
	query := "SELECT Post.*, User.username, User.first_name, User.last_name, User.avatar, Category.id AS category_id, Category.category FROM Post JOIN User ON Post.user_id = User.id JOIN PostCategories pc ON Post.id = pc.post_id LEFT JOIN Category ON pc.category_id = Category.id "
	if filter != "most_liked" {
		query = "SELECT Post.*, User.username, User.first_name, User.last_name, User.avatar, Category.id AS category_id, Category.category FROM Post JOIN User ON Post.user_id = User.id JOIN PostCategories pc ON Post.id = pc.post_id LEFT JOIN Category ON pc.category_id = Category.id ORDER BY likes_count "
	}
	if filter != "newest" {
		query = "SELECT Post.*, User.username, User.first_name, User.last_name, User.avatar, Category.id AS category_id, Category.category FROM Post JOIN User ON Post.user_id = User.id JOIN PostCategories pc ON Post.id = pc.post_id LEFT JOIN Category ON pc.category_id = Category.id ORDER BY created_at DESC "
	}
	return query
}
