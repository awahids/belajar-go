package v1

import (
	"net/http"

	"github.com/awahids/belajar-gin/internal/configs"
	"github.com/awahids/belajar-gin/internal/domain/delivery/handler/bookHandler"
	"github.com/awahids/belajar-gin/internal/domain/infrastructure/db"
	"github.com/awahids/belajar-gin/internal/domain/infrastructure/repositories/bookRepo"
	"github.com/awahids/belajar-gin/internal/domain/services/bookService"
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

	bookRepos := bookRepo.NewBookRepository(db)
	bookService := bookService.NewBookService(bookRepos, validate)
	bookHandler := bookHandler.NewBookHandler(bookService)

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/helloworld")
		{
			eg.GET("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
			})
		}

		v1.GET("/books", bookHandler.GetBooks)
		book := v1.Group("/book")
		{
			book.POST("", bookHandler.CreateBook)
			book.PUT("", bookHandler.UpdateBook)
			book.GET("/:id", bookHandler.GetBook)
			book.DELETE("/:id", bookHandler.DeleteBook)
		}
	}

	return r
}
