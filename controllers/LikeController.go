package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"golang-rest-api-starter/internals/helpers"
	model "golang-rest-api-starter/models"
	"golang-rest-api-starter/service/crud"
	"log"
	"net/http"
)

type Action struct {
	TableToUpdate          string // Post || Comments : (to updates the likes_count || dislikes_count)
	ActionTableTemp        string // Postlikes || Commentlikes
	OppositActionTableTemp string // Postdislikes || Commentdislikes
	IdentifierInDB         string // post_id || comment_id
	IdentifierInHTML       string // postID ||
	Data                   map[string]interface{}
}

// This is meant to defind what type of action the user make.
func (a *Action) Init(r *http.Request, actionTable, oppositActionTable string) {
	a.Data = helpers.GetFormData(r)
	// here make scalable the like action so the user either like a post or a comment.
	a.ActionTableTemp, a.OppositActionTableTemp = fmt.Sprintf("Post%s", actionTable), fmt.Sprintf("Post%s", oppositActionTable)
	a.IdentifierInDB, a.IdentifierInHTML = "post_id", "postID"
	a.TableToUpdate = "Post"
	var _, isIdentifierHTMLIsApartOfForm = a.Data[a.IdentifierInHTML]
	if !isIdentifierHTMLIsApartOfForm {
		a.ActionTableTemp, a.OppositActionTableTemp = fmt.Sprintf("Comment%s", actionTable), fmt.Sprintf("Comment%s", oppositActionTable)
		a.IdentifierInDB, a.IdentifierInHTML = "comment_id", "commentID"
		a.TableToUpdate = "Comments"
	}
}

func (a Action) RemoveALikesOrDislikes(db *sql.DB, response *model.Reponse, table, userID string) {
	if _, errdelete := db.Exec(fmt.Sprintf(`DELETE FROM %s WHERE user_id = ? AND %s = ?`, table, a.IdentifierInDB), userID, a.Data[a.IdentifierInHTML]); errdelete != nil {
		log.Println(errdelete)
		helpers.ErrorWriter(response, "Cannot perform this operation. Please try again", http.StatusInternalServerError)
	}
}

func (a Action) UpdateLikesActionCountInPostOrComment(db *sql.DB, response *model.Reponse, operationSign, table string) {
	updateQuery := fmt.Sprintf("UPDATE %s SET %s_count = %s_count %s ? WHERE id = ?", a.TableToUpdate, table, table, operationSign)
	var _, err = db.Exec(updateQuery, 1, a.Data[a.IdentifierInHTML])
	if err != nil {
		log.Println(err)
		helpers.ErrorWriter(response, "Cannot perform this operation. Please try again", http.StatusInternalServerError)
	}
}

func (a *Action) Create(r *http.Request, db *sql.DB, response *model.Reponse, actionTable, oppositActionTable string) {

	a.Init(r, actionTable, oppositActionTable)
	var err = helpers.IsRequiredFeildsExits(a.Data, a.IdentifierInHTML)
	if err != nil {
		helpers.ErrorWriter(response, err.Error(), http.StatusBadRequest)
	}
	var userID, _ = r.Context().Value("userID").(string)
	if userID == "" {
		helpers.ErrorWriter(response, "You don't have authorization to perfom this action", 403)
	}
	conditions := map[string]interface{}{
		"user_id":        userID,
		a.IdentifierInDB: a.Data[a.IdentifierInHTML],
	}
	columns := []string{"id"}
	var id int

	if SelectErr := crud.Select(db, a.ActionTableTemp, "", conditions, columns, &id); SelectErr != nil && errors.Is(SelectErr, sql.ErrNoRows) {
		columnsToInsert := []string{a.IdentifierInDB, "user_id"}
		if _, InsertionErr := crud.Insert(db, a.ActionTableTemp, columnsToInsert, a.Data[a.IdentifierInHTML], userID); InsertionErr != nil {
			log.Println(InsertionErr)
			helpers.ErrorWriter(response, "Cannot like the post. Please try again", http.StatusInternalServerError)
		}
		if response.StatusCode == http.StatusOK {
			if SelectErr := crud.Select(db, a.OppositActionTableTemp, "", conditions, columns, &id); !(SelectErr != nil && errors.Is(SelectErr, sql.ErrNoRows)) {
				a.UpdateLikesActionCountInPostOrComment(db, response, "-", oppositActionTable)
			}

			a.UpdateLikesActionCountInPostOrComment(db, response, "+", actionTable)
			a.RemoveALikesOrDislikes(db, response, a.OppositActionTableTemp, userID)
		}
	} else {
		a.UpdateLikesActionCountInPostOrComment(db, response, "-", actionTable)
		a.RemoveALikesOrDislikes(db, response, a.ActionTableTemp, userID)
	}
}
