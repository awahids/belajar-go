package auth_service

import (
	"errors"

	"github.com/awahids/belajar-go/internal/delivery/data/dtos"
	"github.com/awahids/belajar-go/internal/domain/models"
	"github.com/awahids/belajar-go/internal/domain/repositories/user_repository"
	"github.com/awahids/belajar-go/pkg/utils"
)

type AuthService struct {
	repo user_repository.UserInterface
}

func NewAuthService(repo user_repository.UserInterface) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Login(loginReq *dtos.LoginReq) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(loginReq.Email)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(loginReq.Password, user.Password) {
		return nil, errors.New("incorrect password")
	}

	return user, nil
}
