package handlers

import (
	"html/template"
	"net/http"
)

type PasswordReset struct {
	template *template.Template
}

func NewPasswordReset(templatePath string) (*PasswordReset, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return nil, err
	}

	return &PasswordReset{
		template: tmpl,
	}, nil
}

func (h *PasswordReset) PasswordResetPage(w http.ResponseWriter, r *http.Request) {
	// Render the login page template
	err := h.template.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}