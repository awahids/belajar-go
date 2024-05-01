package bookService

import (
	"github.com/awahids/belajar-gin/internal/domain/infrastructure/repositories/repoInterface"
	"github.com/awahids/belajar-gin/internal/domain/models"
	serviceinterface "github.com/awahids/belajar-gin/internal/domain/services/serviceInterface"
	"github.com/awahids/belajar-gin/pkg/helpers"
	"github.com/awahids/belajar-gin/pkg/helpers/request/bookReq"
	"github.com/awahids/belajar-gin/pkg/helpers/response/bookRes"
	"github.com/go-playground/validator/v10"
)

type BookService struct {
	repo     repoInterface.BookInterface
	Validate *validator.Validate
}

func NewBookService(repoInterface repoInterface.BookInterface, validate *validator.Validate) serviceinterface.BookService {
	return &BookService{
		repo:     repoInterface,
		Validate: validate,
	}
}

func (s *BookService) GetAllBooks() []bookRes.BookRes {
	result := s.repo.GetAll()

	var books []bookRes.BookRes
	for _, value := range result {
		book := bookRes.BookRes{
			// Id:     value.Id,
			// UUID:   value.UUID,
			Title:  value.Title,
			Author: value.Author,
			Year:   value.Year,
		}
		books = append(books, book)
	}

	return books
}

func (s *BookService) GetBookById(uuid string) bookRes.BookRes {
	book, err := s.repo.GetByID(uuid)
	helpers.ErrorPanic(err)

	bookRes := bookRes.BookRes{
		// Id:     book.Id,
		// UUID:   book.UUID,
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
	}
	return bookRes
}

func (s *BookService) CreateBook(book bookReq.CreateBookReq) {
	err := s.Validate.Struct(book)
	helpers.ErrorPanic(err)

	bookModel := models.Book{
		Title:  book.Title,
		Author: book.Author,
	}
	s.repo.Create(bookModel)
}

func (s *BookService) UpdateBook(bookReq bookReq.UpdateBookReq) {
	book, err := s.repo.GetByID(bookReq.UUID)
	helpers.ErrorPanic(err)

	book.Title = bookReq.Title
	book.Author = bookReq.Author
	book.Year = bookReq.Year
	s.repo.Update(book)
}

func (s *BookService) DeleteBook(uuid string) error {
	return s.repo.Delete(uuid)
}
