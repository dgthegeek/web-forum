package handlers

import (
	"forum/internal/services"
	"html/template"
	"log"
	"net/http"
)

type HomeHandler struct {
	template     *template.Template
	postServices services.PostService
}

func NewHomeHandler(templatePath string, postServices services.PostService) (*HomeHandler, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return nil, err
	}

	return &HomeHandler{
		template:     tmpl,
		postServices: postServices,
	}, nil
}

func (h *HomeHandler) HomePage(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD

=======
>>>>>>> master
	posts, err1 := h.postServices.GetAllPosts()
	if err1 != nil {
		log.Printf("Error getting all posts: %v", err1)
		http.Error(w, "Failed to get posts", http.StatusInternalServerError)
		return
	}

	// Render the template with the posts
	err := h.template.Execute(w, posts)
	if err != nil {
		log.Printf("Failed to render template: %v", err)
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}
