package company_service

import (
	"github.com/awahids/belajar-go/common"
	"github.com/awahids/belajar-go/internal/delivery/data/dtos"
	"github.com/awahids/belajar-go/internal/delivery/data/response"
)

type CompanyInterface interface {
	GetAll(pagination *common.Pagination) (companies []*response.CompanyRespone, totalItems int, err error)
	Create(company *dtos.CreateCompanyReq) (*response.CompanyRespone, error)
	GetById(uuid string) (company *response.CompanyRespone, err error)
}
