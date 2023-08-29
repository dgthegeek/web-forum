package repositories

import (
	"database/sql"
)

type SQLiteLikeRepository struct {
	db *sql.DB
}

// Function to increment the number of likes
func (r *SQLitePostRepository) IncrementLikes(postID, userID int) error {

	// Lets check is the user disliked the post
	var dislikeCount int
	err := r.db.QueryRow("SELECT COUNT(*) FROM postDislikes WHERE userID = ? AND postID = ?", userID, postID).Scan(&dislikeCount)
	if err != nil {
		return err
	}
	// 	update the dislike
	if dislikeCount > 0 {
		// User already disliked the post
		var dislike int
		r.db.QueryRow("SELECT dislikes FROM posts WHERE id = ?", postID).Scan(&dislike)

		if dislike > 0 {
			_, err = r.db.Exec("UPDATE posts SET dislikes = dislikes - 1 WHERE id = ?", postID)
			if err != nil {
				return err
			}
			_, err := r.db.Exec("DELETE FROM postDislikes WHERE postID = ? AND userID = ?", postID, userID)
			if err != nil {
				return err
			}

		}

	}
	// And after that we update the like also

	// Check if the user already liked the post
	var likeCount int
	err = r.db.QueryRow("SELECT COUNT(*) FROM postLikes WHERE userID = ? AND postID = ?", userID, postID).Scan(&likeCount)
	if err != nil {
		return err
	}

	if likeCount > 0 {
		// User already liked the post
		return nil
	}

	// Increment the number of likes
	_, err = r.db.Exec("UPDATE posts SET likes = likes + 1 WHERE id = ?", postID)
	if err != nil {
		return err
	}

	// Save the like into the DB
	_, err = r.db.Exec("INSERT INTO postLikes (userID, postID) VALUES (?, ?)", userID, postID)
	if err != nil {
		return err
	}

	return nil
}

// Function to increment the number of dislikes
func (r *SQLitePostRepository) IncrementDislikes(postID, userID int) error {

	// Lets check is the user liked the post
	var likeCount int
	err := r.db.QueryRow("SELECT COUNT(*) FROM postLikes WHERE userID = ? AND postID = ?", userID, postID).Scan(&likeCount)
	if err != nil {
		return err
	}
	// 	u	pdate the dislike
	if likeCount > 0 {
		// User already liked the post
		var likes int
		r.db.QueryRow("SELECT likes FROM posts WHERE id = ?", postID).Scan(&likes)
		if likes > 0 {
			_, err = r.db.Exec("UPDATE posts SET likes = likes - 1 WHERE id = ?", postID)
			if err != nil {
				return err
			}
			_, err := r.db.Exec("DELETE FROM postLikes WHERE postID = ? AND userID = ?", postID, userID)
			if err != nil {
				return err
			}
		}

	}

	// Check if the user already disliked the post
	var dislikeCount int
	err = r.db.QueryRow("SELECT COUNT(*) FROM postDislikes WHERE userID = ? AND postID = ?", userID, postID).Scan(&dislikeCount)
	if err != nil {
		return err
	}

	if dislikeCount > 0 {
		// User already liked the post
		return nil
	}

	_, err = r.db.Exec("UPDATE posts SET dislikes = dislikes + 1 WHERE id = ?", postID)
	if err != nil {
		return err
	}

	// NOW LETS SAVE THE LIKE INTO THE DB
	r.db.Exec("INSERT INTO postDislikes (userID, postID) VALUES (?, ?)", userID, postID)

	return nil
}
