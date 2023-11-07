package handlers

import (
	"errors"
	internals "golang-rest-api-starter/internals/config/database"
	"golang-rest-api-starter/internals/helpers"
	model "golang-rest-api-starter/models"
	"golang-rest-api-starter/service/crud"
	"log"
	"net/http"

	"github.com/mattn/go-sqlite3"
)

type Auth struct{}

func (a *Auth) Login(w http.ResponseWriter, r *http.Request) {
	db, errInstance := r.Context().Value("db").(model.DB)
	session := Session{}

	internals.TablesCreation(db.Instance)

	// if got error getting database instance.
	if !errInstance {
		helpers.ErrorThrower(errors.New("Error getting database instance"), "Something went wrong", http.StatusInternalServerError, w, r)
		return
	}

	response := model.Reponse{
		Message:    "OK",
		StatusCode: http.StatusOK,
	}
	if r.Method == http.MethodPost {
		datas := helpers.GetFormData(r)
		//server side validation 🔒
		helpers.ValidateForm(datas, &response)
		var err = helpers.IsRequiredFeildsExits(datas, "email", "password")
		if err != nil {
			helpers.ErrorWriter(&response, err.Error(), http.StatusBadRequest)
		}
		//
		if response.StatusCode == http.StatusOK {
			columns := []string{}
			conditions := map[string]interface{}{
				"email": datas["email"],
			}

			var usr model.User
			SelectionErr := crud.Select(db.Instance, "User", "", conditions, columns, &usr.ID, &usr.FirstName, &usr.LastName, &usr.Email, &usr.Username, &usr.Bio, &usr.Avatar, &usr.Password, &usr.CreationDate)

			if SelectionErr != nil || !helpers.PasswordDecrypter(usr.Password, datas["password"].(string)) {
				helpers.ErrorWriter(&response, "Invalid email or password", 401)
			}

			if response.StatusCode == http.StatusOK {
				sessionErr := session.CreateSession(db.Instance, usr, w)
				if sessionErr != nil {
					log.Println(sessionErr)
					helpers.ErrorWriter(&response, "*Failed to login. Try again later", http.StatusBadRequest)
				}
				if response.StatusCode == http.StatusOK {
					response.StatusCode = http.StatusFound

					http.Redirect(w, r, "/posts", response.StatusCode)
					return
				}
			}
		}
	}

	helpers.ResponseFormatter(response, "login", w, r, response.StatusCode)
}

func (a *Auth) Register(w http.ResponseWriter, r *http.Request) {
	db, _ := r.Context().Value("db").(model.DB)
	// Get the submitted data from the client
	datas := helpers.GetFormData(r)
	// Default response sent by the server
	response := model.Reponse{
		Message:    "OK",
		StatusCode: http.StatusOK,
		HasError:   true,
	}
	if r.Method == http.MethodPost {
		//server side validation 🔒
		helpers.ValidateForm(datas, &response)
		var err = helpers.IsRequiredFeildsExits(datas, "fname", "lname", "email", "username", "password")
		if err != nil {
			helpers.ErrorWriter(&response, err.Error(), http.StatusBadRequest)
		}
		bio, ok := datas["bio"]
		if ok {
			if len(bio.(string)) > 255 {
				helpers.ErrorWriter(&response, "Your bio cannot exceed 255 characters.", http.StatusBadRequest)
			}
		}
		// if datas are valid ✅
		if response.StatusCode == http.StatusOK {
			columnsToInsert := []string{"first_name", "last_name", "email", "username", "bio", "avatar", "password"}
			randomAvatar := helpers.AvatarGenerator()

			hashedPassword := helpers.PasswordEncrypter(datas["password"].(string))

			if _, InsertionErr := crud.Insert(db.Instance, "user", columnsToInsert, datas["fname"], datas["lname"], datas["email"], datas["username"], bio, randomAvatar, hashedPassword); InsertionErr != nil {
				log.Println(InsertionErr)
				if sqliteErr, ok := InsertionErr.(sqlite3.Error); ok && sqliteErr.Code == sqlite3.ErrConstraint {
					helpers.ErrorWriter(&response, "*This email or username already exists. Please use a different one", http.StatusConflict)
				} else {
					helpers.ErrorWriter(&response, "Something went wrong if the problem persists, please contact the us", http.StatusInternalServerError)
				}
			}
			if response.StatusCode == http.StatusOK {
				response.StatusCode = 201
				response.HasError = false
			}
		}
	}
	// Send response to the client
	helpers.ResponseFormatter(response, "register", w, r, response.StatusCode)
}

func (a *Auth) Logout(w http.ResponseWriter, r *http.Request) {
	db, _ := r.Context().Value("db").(model.DB)
	userID, _ := r.Context().Value("userID").(string)
	db.Instance.Exec("DELETE FROM Sessions WHERE token = ?", userID)
}
