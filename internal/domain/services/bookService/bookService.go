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

func (s *BookService) GetAllBooks() ([]bookRes.BookRes, error) {
	result, err := s.repo.GetAll()
	helpers.ErrorPanic(err)

	var books []bookRes.BookRes
	for _, value := range result {
		book := bookRes.BookRes{
			Id:     int(value.Id),
			UUID:   value.UUID,
			Title:  value.Title,
			Author: value.Author,
			Year:   value.Year,
		}
		books = append(books, book)
	}

	return books, nil
}

func (s *BookService) GetBookById(uuid string) (bookRes.BookRes, error) {
	book, err := s.repo.GetByID(uuid)
	helpers.ErrorPanic(err)

	bookRes := bookRes.BookRes{
		Id:     int(book.Id),
		UUID:   book.UUID,
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
	}
	return bookRes, nil
}

func (s *BookService) CreateBook(book bookReq.CreateBookReq) (bookcreate bookRes.BookRes, err error) {
	validator := s.Validate.Struct(book)
	helpers.ErrorPanic(validator)

	bookModel := models.Book{
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
	}
	createdBook, err := s.repo.Create(bookModel)
	if err != nil {
		return bookcreate, err
	}

	bookcreate = bookRes.BookRes{
		Id:     int(createdBook.Id),
		UUID:   createdBook.UUID,
		Title:  createdBook.Title,
		Author: createdBook.Author,
		Year:   createdBook.Year,
	}
	return bookcreate, nil
}

func (s *BookService) UpdateBook(bookReq bookReq.UpdateBookReq) bookRes.BookRes {
	book, err := s.repo.GetByID(bookReq.UUID)
	helpers.ErrorPanic(err)

	book.Title = bookReq.Title
	book.Author = bookReq.Author
	book.Year = bookReq.Year
	s.repo.Update(book)

	bookRes := bookRes.BookRes{
		Id:     int(book.Id),
		UUID:   book.UUID,
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
	}

	return bookRes
}

func (s *BookService) DeleteBook(uuid string) error {
	return s.repo.Delete(uuid)
}
