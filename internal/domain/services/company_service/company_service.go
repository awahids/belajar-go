package company_service

import (
	"strings"

	"github.com/awahids/belajar-go/common"
	"github.com/awahids/belajar-go/internal/delivery/data/dtos"
	"github.com/awahids/belajar-go/internal/delivery/data/response"
	"github.com/awahids/belajar-go/internal/domain/models/companies"
	"github.com/awahids/belajar-go/internal/domain/repositories/company_repository"
	"github.com/awahids/belajar-go/pkg/helpers"
	"github.com/awahids/belajar-go/pkg/utils"
	"github.com/go-playground/validator/v10"
)

type CompanyService struct {
	companyRepo company_repository.CompanyInterface
	Validate    *validator.Validate
}

func NewCompanyService(companyRepo company_repository.CompanyInterface, validate *validator.Validate) *CompanyService {
	return &CompanyService{
		companyRepo: companyRepo,
		Validate:    validate,
	}
}

func (s *CompanyService) Create(company *dtos.CreateCompanyReq) (companyRes *response.CompanyRespone, err error) {
	validator := s.Validate.Struct(company)
	if validator != nil {
		return nil, helpers.ErrorValidator(validator)
	}

	name := strings.ToUpper(company.Name)
	generateCode := utils.GenerateRandomCode(6) + name[:3]

	companyModel := companies.Companies{
		Name: name,
		Code: generateCode,
	}

	createdCompany, err := s.companyRepo.Create(&companyModel)
	if err != nil {
		return nil, err
	}

	companyRes = &response.CompanyRespone{
		Id:   int(createdCompany.Id),
		UUID: createdCompany.UUID,
		Name: createdCompany.Name,
		Code: createdCompany.Code,
	}

	return companyRes, nil
}

func (s *CompanyService) GetAll(pagination *common.Pagination) (companies []*response.CompanyRespone, totalItems int, err error) {
	userModel, totalItems64, err := s.companyRepo.GetAll(pagination)
	if err != nil {
		return nil, 0, err
	}

	totalItems = int(totalItems64)

	companies = []*response.CompanyRespone{}
	for _, user := range userModel {
		companies = append(companies, &response.CompanyRespone{
			Id:   int(user.Id),
			UUID: user.UUID,
			Name: user.Name,
			Code: user.Code,
		})
	}

	return companies, totalItems, nil
}

func (s *CompanyService) GetById(uuid string) (company *response.CompanyRespone, err error) {
	companyModel, err := s.companyRepo.GetById(uuid)
	if err != nil {
		return nil, err
	}

	company = &response.CompanyRespone{
		Id:   int(companyModel.Id),
		UUID: companyModel.UUID,
		Name: companyModel.Name,
		Code: companyModel.Code,
	}

	return company, nil
}
