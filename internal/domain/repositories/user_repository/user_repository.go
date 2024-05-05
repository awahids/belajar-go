package user_repository

import (
	"github.com/awahids/belajar-gin/internal/domain/models"
	"github.com/awahids/belajar-gin/pkg/helpers"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserInterface {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) (userCreate *models.User, err error) {
	err = r.db.Create(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetUser(uuid string) (*models.User, error) {
	var user models.User

	if err := r.db.Where("uuid = ?", uuid).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, helpers.ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByUsername(username string) (user *models.User, err error) {
	user = &models.User{}

	err = r.db.Where("username = ?", username).First(user).Error
	helpers.ErrorPanic(err)

	return user, nil
}

func (r *userRepository) GetUserByEmail(email string) (user *models.User, err error) {
	user = &models.User{}

	err = r.db.Where("email = ?", email).First(user).Error
	helpers.ErrorPanic(err)

	return user, nil
}

func (r *userRepository) GetUsers() (users []models.User, err error) {
	err = r.db.Find(&users).Error
	helpers.ErrorPanic(err)

	return users, nil
}

func (r *userRepository) Update(user *models.User) (userUpdate *models.User, err error) {
	err = r.db.Save(user).Error
	helpers.ErrorPanic(err)

	return user, nil
}

func (r *userRepository) Delete(uuid string) (err error) {
	err = r.db.Delete(&models.User{}, "uuid = ?", uuid).Error
	helpers.ErrorPanic(err)

	return nil
}
