package handlers

import (
	"net/http"

	"github.com/awahids/belajar-go/common"
	"github.com/awahids/belajar-go/internal/delivery/data/dtos"
	"github.com/awahids/belajar-go/internal/delivery/data/response"
	"github.com/awahids/belajar-go/internal/domain/services/company_service"
	"github.com/awahids/belajar-go/pkg/helpers"
	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	companyService company_service.CompanyInterface
}

func NewCompanyHandler(companyInterface company_service.CompanyInterface) *CompanyHandler {
	return &CompanyHandler{
		companyService: companyInterface,
	}
}

func (h *CompanyHandler) GetCompanies(ctx *gin.Context) {
	per_page := 5
	page := 1

	pp, p := common.ExtractPaginationParams(ctx, per_page, page)
	offset := (p - 1) * pp

	pagination := &common.Pagination{
		Limit:  pp,
		Offset: offset,
		Page:   page,
	}

	companies, totalItems, err := h.companyService.GetAll(pagination)
	helpers.ErrorPanic(err)

	meta := common.NewMeta(totalItems, pp, p, offset, len(companies))

	webResponse := response.Response{
		Code:    http.StatusOK,
		Message: "Data Found",
		Data:    companies,
		Meta:    meta,
	}
	helpers.JSONResponse(ctx, webResponse)
}

func (h *CompanyHandler) CreateCompany(ctx *gin.Context) {
	createCompanyReq := dtos.CreateCompanyReq{}
	err := ctx.ShouldBindJSON(&createCompanyReq)
	helpers.ErrorPanic(err)

	createdCompany, err := h.companyService.Create(&createCompanyReq)
	if err != nil {
		helpers.ErrorResponse(err, ctx)
		return
	}

	webResponse := response.Response{
		Code:    http.StatusCreated,
		Message: "Success Created",
		Data:    createdCompany,
	}
	helpers.JSONResponse(ctx, webResponse)
}
