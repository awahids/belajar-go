package book_service

import (
	"github.com/awahids/belajar-go/internal/delivery/data/dtos"
	"github.com/awahids/belajar-go/internal/delivery/data/response"
)

type BookInterface interface {
	CreateBook(book *dtos.CreateBookReq) (bookRes *response.BookResponse, err error)
	UpdateBook(book *dtos.UpdateBookReq) (bookRes *response.BookResponse, err error)
	GetBookById(uuid string) (book *response.BookResponse, err error)
	GetAllBooks() (books []*response.BookResponse, err error)
	DeleteBook(uuid string) (err error)
}
