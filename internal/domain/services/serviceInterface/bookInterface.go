package serviceInterface

import (
	"github.com/awahids/belajar-gin/pkg/helpers/request/bookReq"
	"github.com/awahids/belajar-gin/pkg/helpers/response/bookRes"
)

type BookService interface {
	CreateBook(book bookReq.CreateBookReq)
	UpdateBook(book bookReq.UpdateBookReq)
	GetBookById(uuid string) bookRes.BookRes
	DeleteBook(uuid string) error
	GetAllBooks() []bookRes.BookRes
}
