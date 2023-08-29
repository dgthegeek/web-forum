// handlers/post_handler.go

package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"forum/internal/models"
	"forum/internal/services"
)

type PostHandler struct {
	postService services.PostService
	template    *template.Template
}

func NewPostHandler(postService services.PostService, templatePath string) *PostHandler {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return nil
	}
	return &PostHandler{
		postService: postService,
		template:    tmpl,
	}
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var post models.Post

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Printf("Failed to read request body: %v", err)
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	err = h.postService.CreatePost(post)
	if err != nil {
		log.Printf("Failed to create post: %v", err)
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(post)
	if err != nil {
		log.Printf("Failed to marshal post JSON: %v", err)
		http.Error(w, "Failed to marshal post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (h *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	posts, err := h.postService.GetAllPosts()
	if err != nil {
		log.Printf("Failed to get posts: %v", err)
		http.Error(w, "Failed to get posts", http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(posts)
	if err != nil {
		log.Printf("Failed to marshal posts JSON: %v", err)
		http.Error(w, "Failed to marshal posts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (h *PostHandler) LikePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	postID, err := strconv.Atoi(r.URL.Query().Get("id_post"))
	if err != nil {
		log.Printf("Error retrieving the postID: %v", err)
		http.Error(w, "Failed to get data", http.StatusInternalServerError)
		return
	}

<<<<<<< HEAD
	session, _ := session.Storee.Get(r, "currentUser")
=======
	userIDStr, _ := r.Cookie("userID")
	userID, _ := strconv.Atoi(userIDStr.Value)
	log.Println(userID)
>>>>>>> master

	err = h.postService.LikePost(userID, postID)
	if err != nil {
		log.Printf("Error liking the post: %v", err)
		http.Error(w, "Failed to get data", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *PostHandler) DislikePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	postID, err := strconv.Atoi(r.URL.Query().Get("id_post"))
	if err != nil {
		log.Printf("Error retrieving the postID: %v", err)
		http.Error(w, "Failed to get data", http.StatusInternalServerError)
		return
	}

<<<<<<< HEAD
	session, _ := session.Storee.Get(r, "currentUser")
=======
	userIDStr, _ := r.Cookie("userID")
	userID, _ := strconv.Atoi(userIDStr.Value)
	log.Println(userID)
>>>>>>> master

	err = h.postService.DislikePost(userID, postID)
	if err != nil {
		log.Printf("Error disliking the post: %v", err)
		http.Error(w, "Failed to get data", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *PostHandler) GetPostByCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	categoryID := r.URL.Query().Get("id")
	posts, err := h.postService.GetPostsByCategory(categoryID)
	if err != nil {
		log.Printf("Failed to get posts: %v", err)
		http.Error(w, "Failed to get data", http.StatusInternalServerError)
		return
	}

	h.template.Execute(w, posts)
}

func (h *PostHandler) GetPostByUserID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.URL.Query().Get("id")
	posts, err := h.postService.GetPostByUserID(userID)
	if err != nil {
		log.Printf("Failed to get posts: %v", err)
		http.Error(w, "Failed to get data", http.StatusInternalServerError)
		return
	}

	h.template.Execute(w, posts)
}

func (h *PostHandler) GetLikedPostByUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.URL.Query().Get("id")
	posts, err := h.postService.GetLikedPostByUser(userID)
	if err != nil {
		log.Printf("Failed to get posts: %v", err)
		http.Error(w, "Failed to get data", http.StatusInternalServerError)
		return
	}

	h.template.Execute(w, posts)
}
