package auth_service

import (
	"errors"

	"github.com/awahids/belajar-gin/internal/domain/models"
	"github.com/awahids/belajar-gin/internal/domain/repositories/user_repository"
	"github.com/awahids/belajar-gin/pkg/helpers"
)

type AuthService struct {
	repo user_repository.UserInterface
}

func NewAuthService(repo user_repository.UserInterface) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Login(email, password string) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	// Verifikasi password
	if !helpers.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("incorrect password")
	}

	return user, nil
}
