package book_service

import (
	"github.com/awahids/belajar-gin/internal/delivery/data/request"
	"github.com/awahids/belajar-gin/internal/delivery/data/response"
)

type BookInterface interface {
	CreateBook(book *request.CreateBookReq) (bookRes *response.BookResponse, err error)
	UpdateBook(book *request.UpdateBookReq) (bookRes *response.BookResponse, err error)
	GetBookById(uuid string) (book *response.BookResponse, err error)
	GetAllBooks() (books []*response.BookResponse, err error)
	DeleteBook(uuid string) (err error)
}
