package role_repository

import (
	"github.com/awahids/belajar-go/internal/domain/models"
	"github.com/awahids/belajar-go/pkg/helpers"
	"gorm.io/gorm"
)

type RoleRepository struct {
	Db *gorm.DB
}

func NewRoleRepository(Db *gorm.DB) RoleInterface {
	return &RoleRepository{Db: Db}
}

func (r *RoleRepository) GetAll() ([]*models.Roles, error) {
	var roles []*models.Roles
	results := r.Db.Find(&roles)
	helpers.ErrorPanic(results.Error)
	return roles, nil
}

func (r *RoleRepository) GetById(uuid string) (*models.Roles, error) {
	var role *models.Roles
	if err := r.Db.Where("uuid = ?", uuid).First(&role).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, helpers.ErrRecordNotFound
		}
		return nil, err
	}
	return role, nil
}

func (r *RoleRepository) GetByValue(value string) (*models.Roles, error) {
	var role *models.Roles
	if err := r.Db.Where("value = ?", value).First(&role).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, helpers.ErrRecordNotFound
		}
		return nil, err
	}
	return role, nil
}
