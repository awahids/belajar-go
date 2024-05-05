package router

import (
	"github.com/awahids/belajar-go/internal/delivery/handlers"
	"github.com/awahids/belajar-go/internal/domain/repositories/book_repository"
	"github.com/awahids/belajar-go/internal/domain/services/book_service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func BookRouter(group *gin.RouterGroup, validate *validator.Validate, db *gorm.DB) {
	bookRepository := book_repository.NewBookRepository(db)
	bookService := book_service.NewBookService(bookRepository, validate)
	bookHandler := handlers.NewBookHandler(bookService)

	group.GET("/books", bookHandler.GetBooks)
	book := group.Group("/book")
	{
		book.POST("", bookHandler.CreateBook)
		book.GET("/:uuid", bookHandler.GetBook)
		book.PATCH("/:uuid", bookHandler.UpdateBook)
		book.DELETE("/:uuid", bookHandler.DeleteBook)
	}
}
