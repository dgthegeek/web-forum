package services

import (
	"forum/internal/repositories"
)

type LikeService struct {
	likeRepo repositories.SQLiteLikeRepository
}

func NewLikeService(likeRepo repositories.SQLiteLikeRepository) *LikeService {
	return &LikeService{
		likeRepo: likeRepo,
	}
}

func (s *PostService) LikePost(userID, postID int) error {
	// Increment the post's likes count
	err := s.postRepo.IncrementLikes(postID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostService) DislikePost(userID, postID int) error {
	// Increment the post's dislikes count
	err := s.postRepo.IncrementDislikes(postID, userID)
	if err != nil {
		return err
	}

	return nil
}