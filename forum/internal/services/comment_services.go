// services/comment_service.go

package services

import (
	"forum/internal/models"
	"forum/internal/repositories"
)

type CommentService struct {
	commentRepo repositories.SQLiteCommentRepository
}

func NewCommentService(commentRepo repositories.SQLiteCommentRepository) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
	}
}

func (s *CommentService) CreateComment(comment models.Comment) error {
	err := s.commentRepo.SaveComment(comment)
	if err != nil {
		return err
	}

	return nil
}

func (s *CommentService) GetAllComments() ([]models.Comment, error) {
	comments, err := s.commentRepo.GetAllComments()
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (s *CommentService) LikeComment(idComment, userID int) error {
	// Increment the post's likes count
	err := s.commentRepo.IncrementCommentLikes(idComment, userID)
	if err != nil {
		return err
	}

	return nil
}

func (s *CommentService) DislikeComment(idComment, userID int) error {
	// Increment the post's dislikes count
	err := s.commentRepo.IncrementCommentDislikes(idComment, userID)
	if err != nil {
		return err
	}

	return nil
}
