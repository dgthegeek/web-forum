package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type LoginHandler struct {
	template *template.Template
}

func NewLoginHandler(templatePath string) (*LoginHandler, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return nil, err
	}

	return &LoginHandler{
		template: tmpl,
	}, nil
}

func (h *LoginHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	// Render the login page template
	err := h.template.Execute(w, nil)
	if err != nil {
		log.Printf("Failed to render login page template: %v", err)
		http.Error(w, "Failed to render login page", http.StatusInternalServerError)
		return
	}
}
