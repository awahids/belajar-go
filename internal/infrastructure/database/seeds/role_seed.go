package seeds

import (
	"log"

	"github.com/awahids/belajar-go/internal/domain/models"
	"gorm.io/gorm"
)

func RolesSeed(db *gorm.DB) {
	roles := []models.Roles{
		{
			Title: "admin",
			Value: "admin",
		},
		{
			Title: "user",
			Value: "user",
		},
	}

	for _, role := range roles {
		if err := db.Create(&role).Error; err != nil {
			log.Fatalf("could not seed role: %v", err)
		}
	}

	log.Println("Seeding completed")
}
