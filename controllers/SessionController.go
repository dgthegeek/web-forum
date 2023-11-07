package handlers

import (
	"database/sql"
	"errors"
	"golang-rest-api-starter/internals/helpers"
	model "golang-rest-api-starter/models"
	"golang-rest-api-starter/service/crud"
	"log"
	"net/http"
	"time"
)

type Session struct{}

func (s *Session) CreateSession(db *sql.DB, user model.User, w http.ResponseWriter) error {

	randomUUID := helpers.IDGenerator()

	var session = model.Session{
		Token:          randomUUID.String(),
		UserID:         user.ID,
		ExpirationDate: time.Now().Add(24 * time.Hour),
		CreationDate:   time.Now(),
	}

	columns := []string{"token", "expires", "user_id"}
	conditions := map[string]interface{}{
		"user_id": session.UserID,
	}
	columns = []string{"token", "expires", "user_id"}
	var token string
	SelectErr := crud.Select(db, "Session", "", conditions, columns, &token)
	if errors.Is(SelectErr, sql.ErrNoRows) {
		if _, InsertionErr := crud.Insert(db, "Session", columns, session.Token, session.ExpirationDate, session.UserID); InsertionErr != nil {
			return InsertionErr
		}
	} else {
		columns = []string{"token"}
		if _, updateERR := db.Exec("UPDATE Session SET token = ? WHERE user_id = ?", session.Token, session.UserID); updateERR != nil {
			return updateERR
		}

	}

	cookie := http.Cookie{
		Name:     "option-share-session",
		Value:    session.Token,
		Expires:  session.ExpirationDate,
		Path:     "/",
		Domain:   "localhost",
		HttpOnly: false,                 // Make the cookie accessible only through HTTP(S)
		Secure:   true,                  // Set the Secure flag to ensure the cookie is only sent over HTTPS
		SameSite: http.SameSiteNoneMode, // Update SameSite to None.
	}
	http.SetCookie(w, &cookie)
	return nil
}

func (s *Session) IsSessionExist(db *model.DB, r *http.Request) (bool, string) {
	RemovingExpiredSession(db.Instance)

	sessionToken, Err := GetSessionToken(r)
	if Err != nil {
		return false, ""
	}
	var (
		loggedUserID string
		expires      time.Time
	)

	columns := []string{"user_id", "expires"}
	conditions := map[string]interface{}{
		"token": sessionToken,
	}
	// verify if the session token is on the database
	SelectionErr := crud.Select(db.Instance, "Session", "", conditions, columns, &loggedUserID, &expires)
	if SelectionErr != nil {
		log.Println(SelectionErr)
	}
	isSessionValid := SelectionErr == nil && expires.After(time.Now())
	return isSessionValid, loggedUserID
}

func (s *Session) LogOut(r *http.Request) error {
	var db, _ = r.Context().Value("db").(model.DB)
	sessionToken, _ := GetSessionToken(r)
	var _, err = db.Instance.Exec("DELETE FROM Session WHERE token=?", sessionToken)
	return err
}

// This function has as purpose to decode the the user session token
func GetSessionToken(r *http.Request) (string, error) {
	var userSession, errGettingSession = r.Cookie("option-share-session")

	if errGettingSession != nil {
		return "", errGettingSession
	}
	return userSession.Value, nil
}

func RemovingExpiredSession(db *sql.DB) error {

	var rows, err = db.Query(`SELECT Session.token, Session.expires FROM session`)
	if err != nil {
		return err
	}
	var sessions []model.Session
	for rows.Next() {
		var session model.Session
		err = rows.Scan(&session.Token, &session.ExpirationDate)
		if err != nil {
			return err
		}
		sessions = append(sessions, session)
	}

	for _, session := range sessions {
		if !session.ExpirationDate.After(time.Now()) {
			_, err = db.Exec("DELETE FROM Session WHERE token=?", session.Token)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
