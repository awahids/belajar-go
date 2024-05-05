package database

import (
	"fmt"
	"log"
	"time"

	"github.com/awahids/belajar-go/configs"
	"github.com/awahids/belajar-go/pkg/helpers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	cfg, _ := configs.LoadConfig()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helpers.ErrorPanic(err)

	sqlDB, err := db.DB()
	helpers.ErrorPanic(err)

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Successfully connected to the database")
	MigrateAllModels(db)
	return db, nil
}
