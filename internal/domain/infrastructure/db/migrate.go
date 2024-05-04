package db

import (
	"github.com/awahids/belajar-gin/internal/domain/models"
)

func AutoMigrate() {
	db, _ := NewDB()
	db.AutoMigrate(&models.Book{})
}
