package company_repository

import (
	"github.com/awahids/belajar-go/common"
	"github.com/awahids/belajar-go/internal/domain/models/companies"
	"gorm.io/gorm"
)

type CompanyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) CompanyInterface {
	return &CompanyRepository{db: db}
}

func (r *CompanyRepository) GetAll(pagination *common.Pagination) (companies []*companies.Companies, totalItems int64, err error) {
	err = r.db.Model(&companies).Count(&totalItems).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Order("created_at desc").Limit(pagination.Limit).Offset(pagination.Offset).Find(&companies).Error
	if err != nil {
		return nil, 0, err
	}
	return companies, totalItems, nil
}

func (r *CompanyRepository) Create(company *companies.Companies) (*companies.Companies, error) {
	err := r.db.Create(company).Error
	if err != nil {
		return nil, err
	}
	return company, nil
}

func (r *CompanyRepository) GetById(uuid string) (company *companies.Companies, err error) {
	err = r.db.Where("uuid = ?", uuid).First(&company).Error
	if err != nil {
		return nil, err
	}
	return company, nil
}
