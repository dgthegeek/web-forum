package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"forum/internal/models"
)

type SQLitePostRepository struct {
	db *sql.DB
}

func NewSQLitePostRepository(db *sql.DB) *SQLitePostRepository {
	return &SQLitePostRepository{
		db: db,
	}
}

func (r *SQLitePostRepository) SavePost(post models.Post) error {

	fmt.Println("post data", post)
	result, err3 := r.db.Exec("INSERT INTO posts(title,content,author)  VALUES (?,?,?)",
		post.Title, post.Content, post.Author)
	if err3 != nil {
		return err3
	}

	// Retrieve the ID of the newly inserted post
	postID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// Insert rows into the `post_categories` junction table
	insertPostCategoryQuery := "INSERT INTO post_categories (postID, categoryID) VALUES (?, ?)"
	for _, categoryID := range post.Categories {
		_, err := r.db.Exec(insertPostCategoryQuery, postID, categoryID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *SQLitePostRepository) GetAllPosts() ([]models.Post, error) {
	allPost := []models.Post{}

	rows, err := r.db.Query("SELECT id, title, content, author, likes, dislikes FROM posts")
	if err != nil {
		log.Println("error querying posts:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var content string
		var author string
		var title string
		var likes int
		var dislikes int

		err = rows.Scan(&id, &title, &content, &author, &likes, &dislikes)
		if err != nil {
			log.Println("error query rows", err.Error())
		}

		post := models.Post{
			ID:       strconv.Itoa(id),
			Content:  content,
			Author:   author,
			Title:    title,
			Likes:    likes,
			Dislikes: dislikes,
		}

		allPost = append(allPost, post)
	}
	return allPost, nil
}

func (r *SQLitePostRepository) GetPostsByCategory(categoryID string) ([]models.Post, error) {
	rows, err := r.db.Query(`
		SELECT p.id, p.title, p.content, p.author,p.likes,p.dislikes
		FROM posts p
		INNER JOIN post_categories pc ON p.id = pc.postID
		WHERE pc.categoryID = ?
	`, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []models.Post{}
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author,&post.Likes,&post.Dislikes)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *SQLitePostRepository) GetPostByUserID(userID string) ([]models.Post, error) {

	rows, err := r.db.Query(`
	SELECT p.id, p.title, p.content, p.author,p.likes,p.dislikes
	FROM posts p
	WHERE p.author = ?
	`, userID)

	if err != nil {
		return nil, err
	}

	posts := []models.Post{}
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author,&post.Likes,&post.Dislikes)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *SQLitePostRepository) GetLikedPostByUser(userId string) ([]models.Post, error) {
	rows, err := r.db.Query(`
	SELECT p.id, p.title, p.content, p.author,p.likes,p.dislikes
	FROM posts p
	INNER JOIN postLikes l ON p.id = l.postID
	WHERE l.userID = ?
	`, userId)

	if err != nil {
		return nil, err
	}

	posts := []models.Post{}
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author, &post.Likes, &post.Dislikes)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
