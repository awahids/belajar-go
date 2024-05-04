package bookRouter

import (
	"github.com/awahids/belajar-gin/internal/delivery/handler/bookHandler"
	"github.com/awahids/belajar-gin/internal/domain/repositories/bookRepo"
	"github.com/awahids/belajar-gin/internal/domain/services/bookService"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func BookRouter(group *gin.RouterGroup, validate *validator.Validate, db *gorm.DB) {
	bookRepos := bookRepo.NewBookRepository(db)
	bookServ := bookService.NewBookService(bookRepos, validate)
	bookHand := bookHandler.NewBookHandler(bookServ)

	group.GET("/books", bookHand.GetBooks)
	book := group.Group("/book")
	{
		book.POST("", bookHand.CreateBook)
		book.GET("/:uuid", bookHand.GetBook)
		book.PATCH("/:uuid", bookHand.UpdateBook)
		book.DELETE("/:uuid", bookHand.DeleteBook)
	}
}
