package router

import (
	"net/http"

	"github.com/awahids/belajar-gin/internal/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouters(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	validate := validator.New()

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/helloworld")
		{
			eg.GET("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
			})
		}

		AuthRouter(v1, db)

		v1.Use(middlewares.JWTAuthMiddleware())
		{
			UserRouter(v1, validate, db)
			BookRouter(v1, validate, db)
		}
	}

	return r
}
