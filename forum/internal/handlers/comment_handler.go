package handlers

import (
	"encoding/json"
	"forum/internal/models"
	"forum/internal/services"
	"log"
	"net/http"
)

type CommentHandler struct {
	commentService services.CommentService
}

func NewCommentHandler(commentService services.CommentService) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
	}
}

func (h *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var comment models.Comment

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		log.Printf("Failed to read request body: %v", err)
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	err = h.commentService.CreateComment(comment)
	if err != nil {
		log.Printf("Failed to create comment: %v", err)
		http.Error(w, "Failed to create a comment", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *CommentHandler) GetAllCommentsById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	comments, err := h.commentService.GetAllComments()
	if err != nil {
		log.Printf("Failed to get comments: %v", err)
		http.Error(w, "Failed to get comments", http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(comments)
	if err != nil {
		log.Printf("Failed to marshal comments JSON: %v", err)
		http.Error(w, "Failed to get comments", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (h *CommentHandler) LikesComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var likesInfo models.LikesAndDislikesComment
	err := json.NewDecoder(r.Body).Decode(&likesInfo)
	if err != nil {
		log.Printf("Failed to read request body: %v", err)
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	err = h.commentService.LikeComment(likesInfo.IdComment, likesInfo.IdUser)
	if err != nil {
		log.Printf("Failed to like comment: %v", err)
		http.Error(w, "Failed to like comment", http.StatusInternalServerError)
		return
	}
}

func (h *CommentHandler) DislikesComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var dislikesInfo models.LikesAndDislikesComment
	err := json.NewDecoder(r.Body).Decode(&dislikesInfo)
	if err != nil {
		log.Printf("Failed to read request body: %v", err)
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	err = h.commentService.DislikeComment(dislikesInfo.IdComment, dislikesInfo.IdUser)
	if err != nil {
		log.Printf("Failed to dislike comment: %v", err)
		http.Error(w, "Failed to dislike comment", http.StatusInternalServerError)
		return
	}
}
