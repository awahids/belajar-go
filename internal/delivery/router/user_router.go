package router

import (
	"github.com/awahids/belajar-gin/internal/delivery/handlers"
	"github.com/awahids/belajar-gin/internal/domain/repositories/user_repository"
	"github.com/awahids/belajar-gin/internal/domain/services/user_service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func UserRouter(group *gin.RouterGroup, validate *validator.Validate, db *gorm.DB) {
	userRepository := user_repository.NewUserRepository(db)
	userService := user_service.NewUserService(userRepository, validate)
	userHandler := handlers.NewUserHandler(userService)

	group.GET("/users", userHandler.GetAllUsers)
	user := group.Group("/user")
	{
		user.POST("", userHandler.CreateUser)
		user.GET("/:uuid", userHandler.GetUser)
		user.PATCH("/:uuid", userHandler.UpdateUser)
		user.DELETE("/:uuid", userHandler.DeleteUser)
	}
}
