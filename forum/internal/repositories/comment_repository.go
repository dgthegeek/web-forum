package repositories

import (
	"database/sql"
	"fmt"
	"forum/internal/models"
	"log"
)

type CommentRepository interface {
	SaveComment(*models.Comment) error
	GetCommentsByPostID(postID string) ([]*models.Comment, error)
	GetCommentByID(commentID string) (*models.Comment, error)
	IncrementCommentLikes(idComment, userID int) error
	IncrementCommentDislikes(idComment, userID int) error
}

type SQLiteCommentRepository struct {
	db *sql.DB
}

func NewSqliteCommentRepository(db *sql.DB) *SQLiteCommentRepository {
	return &SQLiteCommentRepository{
		db: db,
	}
}

func (r *SQLiteCommentRepository) SaveComment(comment models.Comment) error {
	fmt.Println(comment.Username, "user")
	_, err3 := r.db.Exec("INSERT INTO comments(commentText,idPost,username)  VALUES (?,?,?)",
		comment.Content, comment.PostID, comment.Username)
	if err3 != nil {
		return err3
	}
	return nil
}

func (r *SQLiteCommentRepository) GetAllComments() ([]models.Comment, error) {
	comments := []models.Comment{}
	rows, err := r.db.Query("SELECT idComment,idPost, commentText ,nbrlikes , nbrdislikes,username FROM comments")
	if err != nil {
		log.Println("error querying comments:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var idComment int
		var commentText string
		var idPost string
		var nbrlikes int
		var nbrdislikes int
		var username string
		err = rows.Scan(&idComment, &idPost, &commentText, &nbrlikes, &nbrdislikes, &username)
		if err != nil {
			log.Println("error query rows", err.Error())
		}

		comment := models.Comment{
			ID:       idComment,
			Content:  commentText,
			PostID:   idPost,
			Likes:    nbrlikes,
			Dislikes: nbrdislikes,
			Username: username,
		}

		comments = append(comments, comment)
	}
	return comments, nil
}

// Function to increment the number of likes
func (r *SQLiteCommentRepository) IncrementCommentLikes(idComment, userID int) error {

	// // Lets check if the user disliked the comment
	var dislikeCount int
	err := r.db.QueryRow("SELECT COUNT(*) FROM commentDislikes WHERE userID = ? AND commentID = ?", userID, idComment).Scan(&dislikeCount)
	if err != nil {
		return err
	}
	// // 	update the dislike
	if dislikeCount > 0 {
		// User already disliked the post
		var dislike int
		r.db.QueryRow("SELECT nbrdislikes FROM comments WHERE idComment = ?", idComment).Scan(&dislike)

		if dislike > 0 {
			_, err = r.db.Exec("UPDATE comments SET nbrdislikes = nbrdislikes - 1 WHERE idComment = ?", idComment)
			if err != nil {
				return err
			}
			_, err := r.db.Exec("DELETE FROM commentDislikes WHERE commentID = ? AND userID = ?", idComment, userID)
			if err != nil {
				return err
			}

		}

	}
	// // And after that we update the like also

	// Check if the user already liked the comment
	var likeCount int
	err = r.db.QueryRow("SELECT COUNT(*) FROM commentLikes WHERE userID = ? AND commentID = ?", userID, idComment).Scan(&likeCount)
	if err != nil {
		return err
	}

	if likeCount > 0 {
		// User already liked the comment
		return nil
	}

	// Increment the number of likes
	_, err = r.db.Exec("UPDATE comments SET nbrlikes = nbrlikes + 1 WHERE idComment = ?", idComment)
	if err != nil {

		return err
	}

	// Save the like into the DB
	_, err = r.db.Exec("INSERT INTO commentLikes (userID, commentID) VALUES (?, ?)", userID, idComment)
	if err != nil {
		return err
	}

	return nil
}

// Function to increment the number of dislikes
func (r *SQLiteCommentRepository) IncrementCommentDislikes(idComment, userID int) error {

	// // Lets check is the user liked the post
	var likeCount int
	err := r.db.QueryRow("SELECT COUNT(*) FROM commentLikes WHERE userID = ? AND commentID = ?", userID, idComment).Scan(&likeCount)
	if err != nil {
		return err
	}
	// update the dislike
	if likeCount > 0 {
		// User already liked the post
		var likes int
		r.db.QueryRow("SELECT nbrlikes FROM comments WHERE idComment = ?", idComment).Scan(&likes)
		if likes > 0 {
			_, err = r.db.Exec("UPDATE comments SET nbrlikes = nbrlikes - 1 WHERE idComment = ?", idComment)
			if err != nil {
				return err
			}
			_, err := r.db.Exec("DELETE FROM commentLikes WHERE commentID = ? AND userID = ?", idComment, userID)
			if err != nil {
				return err
			}
		}

	}

	// // Check if the user already disliked the post
	var dislikeCount int
	err = r.db.QueryRow("SELECT COUNT(*) FROM commentDislikes WHERE userID = ? AND commentID = ?", userID, idComment).Scan(&dislikeCount)
	if err != nil {
		return err
	}

	if dislikeCount > 0 {
		// User already liked the post
		return nil
	}

	_, err = r.db.Exec("UPDATE comments SET nbrdislikes = nbrdislikes + 1 WHERE  idComment = ?", idComment)
	if err != nil {
		return err
	}

	// NOW LETS SAVE THE LIKE INTO THE DB
	r.db.Exec("INSERT INTO commentDislikes (userID, commentID) VALUES (?, ?)", userID, idComment)

	if err != nil {
		return err
	}

	return nil
}
