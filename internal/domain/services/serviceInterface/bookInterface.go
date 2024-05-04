package serviceInterface

import (
	"github.com/awahids/belajar-gin/pkg/helpers/request/bookReq"
	"github.com/awahids/belajar-gin/pkg/helpers/response/bookRes"
)

type BookService interface {
	CreateBook(book bookReq.CreateBookReq) (bookRes bookRes.BookRes, err error)
	UpdateBook(book bookReq.UpdateBookReq) bookRes.BookRes
	GetBookById(uuid string) (book bookRes.BookRes, err error)
	DeleteBook(uuid string) error
	GetAllBooks() (books []bookRes.BookRes, err error)
}
