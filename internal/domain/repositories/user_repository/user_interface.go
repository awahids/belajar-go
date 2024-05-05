package user_repository

import "github.com/awahids/belajar-go/internal/domain/models"

type UserInterface interface {
	Create(user *models.User) (userCreate *models.User, err error)
	GetUser(uuid string) (user *models.User, err error)
	GetByUsername(username string) (user *models.User, err error)
	GetUserByEmail(email string) (user *models.User, err error)
	GetUsers() (users []models.User, err error)
	Update(user *models.User) (userUpdate *models.User, err error)
	Delete(uuid string) (err error)
}
