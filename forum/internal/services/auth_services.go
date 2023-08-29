package services

import (
	"errors"
	"fmt"

	"forum/internal/models"
	"forum/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo repositories.SQLiteUserRepository
}

func NewAuthService(userRepo repositories.SQLiteUserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Login(email, password string) (*models.User, error) {
	user, err := s.userRepo.GetUserByUsername(email)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("USER NOgolang.org/x/crypto/bcryptT FOUND TEST : Invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, errors.New("password Problem test : Invalid email or password")
	}

	return user, nil
}

// services/auth_service.go

// ... (existing code)

func (s *AuthService) ResetPassword(email, newPassword string) error {
	user, err := s.userRepo.GetUserByUsername(email)
	if err != nil {
		return errors.New("User not found")
	}

	err = s.userRepo.UpdatePassword(user)
	if err != nil {
		return err
	}

	return nil
}
