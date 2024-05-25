package router

import (
	"github.com/awahids/belajar-go/internal/delivery/http/handlers"
	"github.com/awahids/belajar-go/internal/domain/repositories/company_repository"
	"github.com/awahids/belajar-go/internal/domain/services/company_service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func CompanyRouter(group *gin.RouterGroup, validate *validator.Validate, db *gorm.DB) {
	companyRepo := company_repository.NewCompanyRepository(db)
	companyService := company_service.NewCompanyService(companyRepo, validate)
	companyHandler := handlers.NewCompanyHandler(companyService)

	group.GET("/companies", companyHandler.GetCompanies)
	book := group.Group("/company")
	{
		book.POST("", companyHandler.CreateCompany)
	}
}
