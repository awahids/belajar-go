package database

import (
	"gorm.io/gorm"
)

func MigrateAllModels(db *gorm.DB) {
	db.AutoMigrate(
	// &models.Book{},
	// &models.User{},
	// &models.Roles{},
	// &companies.Companies{},
	)
}
