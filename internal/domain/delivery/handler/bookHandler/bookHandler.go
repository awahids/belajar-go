package bookHandler

import (
	"net/http"

	"github.com/awahids/belajar-gin/internal/domain/services/serviceInterface"
	"github.com/awahids/belajar-gin/pkg/helpers"
	"github.com/awahids/belajar-gin/pkg/helpers/request/bookReq"
	"github.com/awahids/belajar-gin/pkg/helpers/response"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookService serviceInterface.BookService
}

func NewBookHandler(bookInterface serviceInterface.BookService) *BookHandler {
	return &BookHandler{
		bookService: bookInterface,
	}
}

func (h *BookHandler) GetBooks(ctx *gin.Context) {
	books := h.bookService.GetAllBooks()

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   books,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (h *BookHandler) GetBook(ctx *gin.Context) {
	bookUuid := ctx.Param("uuid")

	bookRes := h.bookService.GetBookById(bookUuid)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   bookRes,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (h *BookHandler) CreateBook(ctx *gin.Context) {
	createBookReq := bookReq.CreateBookReq{}
	err := ctx.ShouldBindJSON(&createBookReq)
	helpers.ErrorPanic(err)

	h.bookService.CreateBook(createBookReq)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (h *BookHandler) UpdateBook(ctx *gin.Context) {
	updateBookReq := bookReq.UpdateBookReq{}
	err := ctx.ShouldBindJSON(&updateBookReq)
	helpers.ErrorPanic(err)

	bookUuid := ctx.Param("uuid")

	updateBookReq.UUID = bookUuid
	h.bookService.UpdateBook(updateBookReq)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (h *BookHandler) DeleteBook(ctx *gin.Context) {
	bookUuid := ctx.Param("uuid")

	h.bookService.DeleteBook(bookUuid)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
