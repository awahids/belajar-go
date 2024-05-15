package user_service

import (
	"errors"

	"github.com/awahids/belajar-go/internal/delivery/data/request"
	"github.com/awahids/belajar-go/internal/delivery/data/response"
	"github.com/awahids/belajar-go/internal/domain/models"
	"github.com/awahids/belajar-go/internal/domain/repositories/role_repository"
	"github.com/awahids/belajar-go/internal/domain/repositories/user_repository"
	"github.com/awahids/belajar-go/pkg/helpers"
	"github.com/go-playground/validator/v10"
)

type UserService struct {
	repo     user_repository.UserInterface
	roleRepo role_repository.RoleInterface
	Validate *validator.Validate
}

func NewUserService(repo user_repository.UserInterface, roleRepo role_repository.RoleInterface, validate *validator.Validate) *UserService {
	return &UserService{
		repo:     repo,
		roleRepo: roleRepo,
		Validate: validate,
	}
}

func (s *UserService) CreateUser(userReq *request.CreateUserReq) (userRes *response.UserResponse, err error) {
	validator := s.Validate.Struct(userReq)
	if validator != nil {
		return nil, helpers.ErrorValidator(validator)
	}

	password, err := helpers.HashPassword(userReq.Password)
	helpers.ErrorPanic(err)

	role, err := s.roleRepo.GetByUuid(userReq.Role.RoleUuid)
	if err != nil {
		return nil, err
	}

	userModel := models.User{
		Username: userReq.Username,
		Email:    userReq.Email,
		Password: password,
		Role:     *role,
	}

	createdUser, err := s.repo.Create(&userModel)
	if err != nil {
		return nil, err
	}

	userRes = &response.UserResponse{
		Id:       int(createdUser.Id),
		UUID:     createdUser.UUID,
		Username: createdUser.Username,
		Email:    createdUser.Email,
		Password: password,
		Role: response.RoleResponse{
			Id:    int(role.Id),
			UUID:  role.UUID,
			Title: string(role.Title),
			Value: role.Value,
		},
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
		Role: response.RoleResponse{
			Id:    int(userModel.Role.Id),
			UUID:  userModel.Role.UUID,
			Title: string(userModel.Role.Title),
			Value: userModel.Role.Value,
		},
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
			Role: response.RoleResponse{
				Id:    int(userModel.Role.Id),
				UUID:  userModel.Role.UUID,
				Title: string(userModel.Role.Title),
				Value: userModel.Role.Value,
			},
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
		Role: response.RoleResponse{
			Id:    int(updatedUser.Role.Id),
			UUID:  updatedUser.Role.UUID,
			Title: string(updatedUser.Role.Title),
			Value: updatedUser.Role.Value,
		},
	}
	return userRes, nil
}

func (s *UserService) DeleteUser(uuid string) error {
	return s.repo.Delete(uuid)
}
