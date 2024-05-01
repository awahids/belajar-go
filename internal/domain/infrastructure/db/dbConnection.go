package db

import (
	"fmt"
	"log"
	"time"

	"github.com/awahids/belajar-gin/internal/configs"
	"github.com/awahids/belajar-gin/internal/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(cfg *configs.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Successfully connected to the database")
	db.AutoMigrate(&models.User{}, &models.Book{})
	return db, nil
}
