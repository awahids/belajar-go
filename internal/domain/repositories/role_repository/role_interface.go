package role_repository

import "github.com/awahids/belajar-go/internal/domain/models"

type RoleInterface interface {
	// GetAll() ([]*models.Roles, error)
	GetByUuid(uuid string) (*models.Roles, error)
	// GetByValue(value string) (*models.Roles, error)
}
