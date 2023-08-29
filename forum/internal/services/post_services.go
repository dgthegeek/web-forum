// services/post_service.go

package services

import (
	"fmt"
	"forum/internal/models"
	"forum/internal/repositories"
)

type PostService struct {
	postRepo repositories.SQLitePostRepository
}

func NewPostService(postRepo repositories.SQLitePostRepository) *PostService {
	return &PostService{
		postRepo: postRepo,
	}
}

func (s *PostService) CreatePost(post models.Post) error {
	err := s.postRepo.SavePost(post)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostService) GetAllPosts() ([]models.Post, error) {
	posts, err := s.postRepo.GetAllPosts()
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *PostService) GetPostByUserID(userID string) ([]models.Post, error) {
	posts, err := s.postRepo.GetPostByUserID(userID)
	if err != nil {
		fmt.Println("sql error: error getting post liked by user", err.Error())
	}
	return posts, nil
}

func (s *PostService) GetPostsByCategory(categoryID string) ([]models.Post, error) {
	posts, err := s.postRepo.GetPostsByCategory(categoryID)
	if err != nil {
		fmt.Println("sql error: error getting post by category", err.Error())
	}
	return posts, nil
}

func (s *PostService) GetLikedPostByUser(userID string) ([]models.Post, error) {
	posts, err := s.postRepo.GetLikedPostByUser(userID)
	if err != nil {
		fmt.Println("sql error: error getting post liked by user", err.Error())
	}
	return posts, nil
}
