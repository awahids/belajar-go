package user_service

import (
	"errors"

	"github.com/awahids/belajar-gin/internal/delivery/data/request"
	"github.com/awahids/belajar-gin/internal/delivery/data/response"
	"github.com/awahids/belajar-gin/internal/domain/models"
	"github.com/awahids/belajar-gin/internal/domain/repositories/user_repository"
	"github.com/awahids/belajar-gin/pkg/helpers"
	"github.com/go-playground/validator/v10"
)

type UserService struct {
	repo     user_repository.UserInterface
	Validate *validator.Validate
}

func NewUserService(userInterface user_repository.UserInterface, validate *validator.Validate) UserInterface {
	return &UserService{
		repo:     userInterface,
		Validate: validate,
	}
}

func (s *UserService) CreateUser(userReq *request.CreateUserReq) (userRes *response.UserResponse, err error) {
	validator := s.Validate.Struct(userReq)
	helpers.ErrorValidator(validator)

	userModel := models.User{}

	password, err := helpers.HashPassword(userReq.Password)
	helpers.ErrorPanic(err)
	userModel.Password = password

	createdUser, err := s.repo.Create(&userModel)
	if err != nil {
		return userRes, err
	}

	userRes = &response.UserResponse{
		Id:       int(createdUser.Id),
		UUID:     createdUser.UUID,
		Username: createdUser.Username,
		Email:    createdUser.Email,
		Password: password,
	}

	return userRes, nil
}

func (s *UserService) GetUserById(uuid string) (user response.UserResponse, err error) {
	userModel, err := s.repo.GetUser(uuid)
	if err != nil {
		if errors.Is(err, helpers.ErrRecordNotFound) {
			return response.UserResponse{}, errors.New("user not found")
		}
		return response.UserResponse{}, err
	}

	userRes := response.UserResponse{
		Id:       int(userModel.Id),
		UUID:     userModel.UUID,
		Username: userModel.Username,
		Email:    userModel.Email,
		Password: userModel.Password,
		// RoleId:   userModel.RoleId,
	}
	return userRes, nil
}

func (s *UserService) GetAllUsers() (users []response.UserResponse, err error) {
	userModels, err := s.repo.GetUsers()
	helpers.ErrorPanic(err)

	for _, userModel := range userModels {
		user := response.UserResponse{
			Id:       int(userModel.Id),
			UUID:     userModel.UUID,
			Username: userModel.Username,
			Email:    userModel.Email,
			Password: userModel.Password,
			// RoleId:   userModel.RoleId,
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *UserService) UpdateUser(userReq request.UpdateUserReq) (userRes response.UserResponse, err error) {
	validator := s.Validate.Struct(userReq)
	helpers.ErrorPanic(validator)

	userModel := models.User{}
	updatedUser, err := s.repo.Update(&userModel)
	if err != nil {
		return userRes, err
	}

	userRes = response.UserResponse{
		Id:       int(updatedUser.Id),
		UUID:     updatedUser.UUID,
		Username: updatedUser.Username,
		Email:    updatedUser.Email,
		Password: updatedUser.Password,
		// RoleId:   updatedUser.RoleId,
	}
	return userRes, nil
}

func (s *UserService) DeleteUser(uuid string) error {
	return s.repo.Delete(uuid)
}
