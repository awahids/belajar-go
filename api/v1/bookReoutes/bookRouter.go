package bookRoutes

import (
	"github.com/awahids/belajar-gin/internal/domain/delivery/handler/bookHandler"
	"github.com/awahids/belajar-gin/internal/domain/infrastructure/db"
	"github.com/awahids/belajar-gin/internal/domain/infrastructure/repositories/bookRepo"
	"github.com/awahids/belajar-gin/internal/domain/services/bookService"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BookRouter(group *gin.RouterGroup, validate *validator.Validate) {
	database, _ := db.NewDB()
	bookRepos := bookRepo.NewBookRepository(database)
	bookServ := bookService.NewBookService(bookRepos, validate)
	bookHand := bookHandler.NewBookHandler(bookServ)

	group.GET("/books", bookHand.GetBooks)
	book := group.Group("/book")
	{
		book.GET("", bookHand.GetBooks)
		book.POST("", bookHand.CreateBook)
		book.GET("/:uuid", bookHand.GetBook)
		book.PATCH("/:uuid", bookHand.UpdateBook)
		book.DELETE("/:uuid", bookHand.DeleteBook)
	}
}
