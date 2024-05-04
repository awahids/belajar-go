package v1

import (
	"net/http"

	bookRoutes "github.com/awahids/belajar-gin/api/v1/bookReoutes"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouters() *gin.Engine {
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

		bookRoutes.BookRouter(v1, validate)
	}

	return r
}
