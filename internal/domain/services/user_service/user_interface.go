package user_service

import (
	"github.com/awahids/belajar-go/internal/delivery/data/dtos"
	"github.com/awahids/belajar-go/internal/delivery/data/response"
)

type UserInterface interface {
	CreateUser(userReq *dtos.CreateUserReq) (userRes *response.UserResponse, err error)
	GetUserById(uuid string) (user response.UserResponse, err error)
	GetAllUsers() (users []response.UserResponse, err error)
	UpdateUser(userReq *dtos.UpdateUserReq) (userRes *response.UserResponse, err error)
	DeleteUser(uuid string) (err error)
}
