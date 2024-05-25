package company_repository

import (
	"github.com/awahids/belajar-go/common"
	companies "github.com/awahids/belajar-go/internal/domain/models/companies"
)

type CompanyInterface interface {
	GetAll(pagination *common.Pagination) (companies []*companies.Companies, items int64, err error)
	Create(company *companies.Companies) (*companies.Companies, error)
	GetById(uuid string) (company *companies.Companies, err error)
	// Update(company *models.Companies) (*models.Companies, error)
	// Delete(uuid string) error
}
