package helpers

import (
	model "golang-rest-api-starter/models"
	"golang-rest-api-starter/service/crud"
	"log"
	"net/http"
	"strings"
	"text/template"
)

/*
This function has as purpose formatted the
response from the server & send it the client
*/
func ResponseFormatter(data any, page string, w http.ResponseWriter, r *http.Request, statusCode int) {
	w.WriteHeader(statusCode)
		log.Println("Request URL:", r.URL.Path, "Status Code:", statusCode)
		
	tmpl := template.Must(template.New("").Funcs(funcMap(r)).ParseGlob("views/*.html"))
	tmpl = template.Must(tmpl.ParseGlob("views/auth/*.html"))
	tmpl = template.Must(tmpl.ParseGlob("views/posts/*.html"))

	tmpl.ExecuteTemplate(w, page+".html", data)
	return
}

func Category(db *model.DB) ([]model.Category, error) {
	var q = "SELECT * FROM Category"
	var rows, err = db.Instance.Query(q)
	if err != nil {
		return nil, err
	}

	var categories []model.Category
	for rows.Next() {
		var category model.Category
		err = rows.Scan(&category.Id, &category.Category)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func funcMap(r *http.Request) template.FuncMap {
	var db, errgettingInstancedb = r.Context().Value("db").(model.DB)
	var userID, errgettinguserID = r.Context().Value("userID").(string)

	var haveAside = func() bool {
		pageThatShouldHaveAside := []string{"/users", "/profile", "/posts"}
		for _, v := range pageThatShouldHaveAside {
			if strings.HasPrefix(r.URL.Path, v) {
				return true
			}
		}
		return false
	}

	return template.FuncMap{
		"isAuthentificated": func() bool {
			var isAuthentificated = false
			if r.Context().Value("isAuthentificated") != nil {
				isAuthentificated = r.Context().Value("isAuthentificated").(bool)
			}
			return isAuthentificated
		},
		"shouldHaveAsides": func() bool {
			return haveAside()
		},
		"Category": func() []model.Category {
			var category []model.Category
			if errgettingInstancedb {
				category, _ = Category(&db)
			}
			return category
		},
		"User": func() model.User {
			var user model.User
			if errgettinguserID && errgettingInstancedb {
				columns := []string{}
				conditions := map[string]interface{}{
					"id": userID,
				}

				if SelectErr := crud.Select(db.Instance, "User", "", conditions, columns, &user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Username, &user.Bio, &user.Avatar, &user.Password, &user.CreationDate); SelectErr != nil {
					log.Println(SelectErr)
				}
			}
			return user
		},
	}
}
