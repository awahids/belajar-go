package v1

import (
	"net/http"

	bookRoutes "github.com/awahids/belajar-gin/api/v1/bookReoutes"
	"github.com/awahids/belajar-gin/internal/configs"
	"github.com/awahids/belajar-gin/internal/domain/infrastructure/db"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouters() *gin.Engine {
	r := gin.Default()

	cfg, _ := configs.LoadConfig()
	db, _ := db.NewDB(cfg)

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

		bookRoutes.BookRouter(v1, validate, db)
	}

	return r
}
