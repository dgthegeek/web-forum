package services

import (
	"errors"

	"forum/internal/models"
	"forum/internal/repositories"
)

type RegistrationService struct {
	userRepo repositories.SQLiteUserRepository
}

func NewRegistrationService(userRepo repositories.SQLiteUserRepository) *RegistrationService {
	return &RegistrationService{
		userRepo: userRepo,
	}
}

func (s *RegistrationService) RegisterUser(user *models.User) error {
	existingUser, err := s.userRepo.GetUserByUsername(user.Username)
	if existingUser != nil || err == nil {
		return errors.New("username already exists")
	}

	err = s.userRepo.SaveUser(user)
	if err != nil {
		return err
	}

	return nil
}
