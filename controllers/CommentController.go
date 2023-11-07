package handlers

import (
	"database/sql"
	"fmt"
	"golang-rest-api-starter/internals/helpers"
	model "golang-rest-api-starter/models"
	"golang-rest-api-starter/service/crud"
	"log"
	"net/http"
)

type Comment struct{}

func (c *Comment) Create(r *http.Request, db *sql.DB, response *model.Reponse) {
	var userID, _ = r.Context().Value("userID").(string)
	if userID == "" {
		helpers.ErrorWriter(response, "You don't have authorization to perfom this action", 403)
	}

	datas := helpers.GetFormData(r)
	helpers.ValidateForm(datas, response)
	var err = helpers.IsRequiredFeildsExits(datas, "postID", "comment")
	if err != nil {
		helpers.ErrorWriter(response, err.Error(), http.StatusBadRequest)
	}
	if len(datas["comment"].(string)) > 255 {
		helpers.ErrorWriter(response, "Your comment cannot exceed 255 characters.", http.StatusBadRequest)
	}
	if response.StatusCode == http.StatusOK {
		columnsToInsert := []string{"post_id", "user_id", "comment"}
		if _, InsertionErr := crud.Insert(db, "Comments", columnsToInsert, datas["postID"], userID, datas["comment"]); InsertionErr != nil {
			log.Println(InsertionErr)
			helpers.ErrorWriter(response, "Failed to post your comment. Try again later", http.StatusInternalServerError)
		} else {
			c.UpdateCommentsCountsInActuallPost(db, response, datas["postID"])
		}
	}
}

func (c *Comment) UpdateCommentsCountsInActuallPost(db *sql.DB, response *model.Reponse, postID interface{}) {
	updateQuery := `UPDATE Post SET comments_count = comments_count + ? WHERE id = ?`
	var _, err = db.Exec(updateQuery, 1, postID)
	if err != nil {
		log.Println(err)
		helpers.ErrorWriter(response, "Cannot perform this operation. Please try again", http.StatusInternalServerError)
	}
}

func (c *Comment) GetPostComments(db *sql.DB, id string, limit string) ([]model.Comment, error) {
	var query = fmt.Sprintf(`SELECT Comments.*, User.username, User.first_name, User.last_name, User.avatar FROM Comments JOIN User ON Comments.user_id = User.id WHERE post_id = ? ORDER BY Comments.created_at DESC %s`, limit)

	// Logic for retreiving post comments
	rows, err := db.Query(query, id)
	if err != nil {
		return []model.Comment{}, err
	}
	defer rows.Close()
	var comments []model.Comment
	for rows.Next() {
		var comment model.Comment
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.AuthorID, &comment.Content, &comment.LikesCount, &comment.DisLikesCount, &comment.CreationDate, &comment.UserName, &comment.FirstName, &comment.LastName, &comment.Avatar)
		if err != nil {
			return []model.Comment{}, err
		}
		// the current user has comment this if we found like.AuthorID == currentUserID.
		// Get the number of like and dislike 💖 relative to the post

		comments = append(comments, comment)
	}
	return comments, nil

}
