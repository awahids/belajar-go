package router

import (
	"github.com/awahids/belajar-gin/internal/delivery/handlers/auth_handler"
	"github.com/awahids/belajar-gin/internal/domain/repositories/user_repository"
	"github.com/awahids/belajar-gin/internal/domain/services/auth_service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRouter(group *gin.RouterGroup, db *gorm.DB) {
	userRepository := user_repository.NewUserRepository(db)
	authService := auth_service.NewAuthService(userRepository)
	authHandler := auth_handler.NewAuthHandler(authService)

	book := group.Group("/auth")
	{
		book.POST("/login", authHandler.Login)
	}
}
