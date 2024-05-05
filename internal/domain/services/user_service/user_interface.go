package user_service

import (
	"github.com/awahids/belajar-gin/internal/delivery/data/request"
	"github.com/awahids/belajar-gin/internal/delivery/data/response"
)

type UserInterface interface {
	CreateUser(userReq *request.CreateUserReq) (userRes *response.UserResponse, err error)
	GetUserById(uuid string) (user response.UserResponse, err error)
	GetAllUsers() (users []response.UserResponse, err error)
	UpdateUser(userReq request.UpdateUserReq) (userRes response.UserResponse, err error)
	DeleteUser(uuid string) (err error)
}
