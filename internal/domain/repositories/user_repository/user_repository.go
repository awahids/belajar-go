package user_repository

import (
	"github.com/awahids/belajar-go/internal/domain/models"
	"github.com/awahids/belajar-go/pkg/helpers"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserInterface {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUser(uuid string) (*models.User, error) {
	var user models.User

	if err := r.db.Preload("Role").Where("uuid = ?", uuid).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, helpers.ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByUsername(username string) (user *models.User, err error) {
	user = &models.User{}

	err = r.db.Where("username = ?", username).First(user).Error
	helpers.ErrorPanic(err)

	return user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (user *models.User, err error) {
	user = &models.User{}

	err = r.db.Where("email = ?", email).First(user).Error
	helpers.ErrorPanic(err)

	return user, nil
}

func (r *UserRepository) GetUsers() (users []models.User, err error) {
	err = r.db.Preload("Role").Find(&users).Error
	helpers.ErrorPanic(err)

	return users, nil
}

func (r *UserRepository) Update(user *models.User) (userUpdate *models.User, err error) {
	err = r.db.Save(user).Error
	helpers.ErrorPanic(err)

	return user, nil
}

func (r *UserRepository) Delete(uuid string) (err error) {
	err = r.db.Delete(&models.User{}, "uuid = ?", uuid).Error
	helpers.ErrorPanic(err)

	return nil
}
