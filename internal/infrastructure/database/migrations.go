package database

import (
	"github.com/awahids/belajar-go/internal/domain/models"
	"gorm.io/gorm"
)

func MigrateAllModels(db *gorm.DB) {
	db.AutoMigrate(
		&models.Book{},
		&models.User{},
	)
}
