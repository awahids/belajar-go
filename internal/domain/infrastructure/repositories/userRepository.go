package repositories

import (
	"database/sql"

	"github.com/awahids/belajar-gin/internal/domain/models"
)

type UserRepository interface {
	FindByUsername(username string) (*models.User, error)
}

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{DB: db}
}

func (repo *userRepository) FindByUsername(username string) (*models.User, error) {
	user := &models.User{}
	err := repo.DB.QueryRow("SELECT id, username, email FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
