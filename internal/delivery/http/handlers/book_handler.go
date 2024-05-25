package handlers

import (
	"net/http"

	"github.com/awahids/belajar-go/internal/delivery/data/dtos"
	"github.com/awahids/belajar-go/internal/delivery/data/response"
	"github.com/awahids/belajar-go/internal/domain/services/book_service"
	"github.com/awahids/belajar-go/pkg/helpers"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookService book_service.BookInterface
}

func NewBookHandler(bookInterface book_service.BookInterface) *BookHandler {
	return &BookHandler{
		bookService: bookInterface,
	}
}

func (h *BookHandler) GetBooks(ctx *gin.Context) {
	books, err := h.bookService.GetAllBooks()
	helpers.ErrorPanic(err)

	webResponse := response.Response{
		Code:    http.StatusOK,
		Message: "Data Found",
		Data:    books,
	}

	helpers.JSONResponse(ctx, webResponse)
}

func (h *BookHandler) GetBook(ctx *gin.Context) {
	bookUuid := ctx.Param("uuid")

	bookRes, err := h.bookService.GetBookById(bookUuid)
	helpers.ErrorPanic(err)

	webResponse := response.Response{
		Code:    http.StatusOK,
		Message: "Data Found",
		Data:    bookRes,
	}

	helpers.JSONResponse(ctx, webResponse)
}

func (h *BookHandler) CreateBook(ctx *gin.Context) {
	createBookReq := dtos.CreateBookReq{}
	err := ctx.ShouldBindJSON(&createBookReq)
	helpers.ErrorPanic(err)

	createdBook, err := h.bookService.CreateBook(&createBookReq)
	helpers.ErrorPanic(err)

	webResponse := response.Response{
		Code:    http.StatusCreated,
		Message: "Success Created",
		Data:    createdBook,
	}

	helpers.JSONResponse(ctx, webResponse)
}

func (h *BookHandler) UpdateBook(ctx *gin.Context) {
	updateBookReq := dtos.UpdateBookReq{}
	err := ctx.ShouldBindJSON(&updateBookReq)
	helpers.ErrorPanic(err)

	bookUUID := ctx.Param("uuid")
	updateBookReq.UUID = bookUUID

	updatedBook, err := h.bookService.UpdateBook(&updateBookReq)
	helpers.ErrorPanic(err)

	webResponse := response.Response{
		Code:    http.StatusOK,
		Message: "Success Updated",
		Data:    updatedBook,
	}

	helpers.JSONResponse(ctx, webResponse)
}

func (h *BookHandler) DeleteBook(ctx *gin.Context) {
	bookUuid := ctx.Param("uuid")

	h.bookService.DeleteBook(bookUuid)

	webResponse := response.Response{
		Code:    http.StatusOK,
		Message: "Success Deleted",
		Data:    nil,
	}

	helpers.JSONResponse(ctx, webResponse)
}
