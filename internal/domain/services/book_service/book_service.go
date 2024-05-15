package book_service

import (
	"github.com/awahids/belajar-go/internal/delivery/data/request"
	"github.com/awahids/belajar-go/internal/delivery/data/response"
	"github.com/awahids/belajar-go/internal/domain/models"
	"github.com/awahids/belajar-go/internal/domain/repositories/book_repository"
	"github.com/awahids/belajar-go/pkg/helpers"
	"github.com/go-playground/validator/v10"
)

type BookService struct {
	repo     book_repository.BookInterface
	Validate *validator.Validate
}

func NewBookService(bookInterface book_repository.BookInterface, validate *validator.Validate) BookInterface {
	return &BookService{
		repo:     bookInterface,
		Validate: validate,
	}
}

func (s *BookService) GetAllBooks() ([]*response.BookResponse, error) {
	result, err := s.repo.GetAll()
	helpers.ErrorPanic(err)

	var books []*response.BookResponse
	for _, value := range result {
		book := response.BookResponse{
			Id:     int(value.Id),
			UUID:   value.UUID,
			Title:  value.Title,
			Author: value.Author,
			Year:   value.Year,
		}
		books = append(books, &book)
	}

	return books, nil
}

func (s *BookService) GetBookById(uuid string) (*response.BookResponse, error) {
	book, err := s.repo.GetByID(uuid)
	helpers.ErrorPanic(err)

	bookRes := &response.BookResponse{
		Id:     int(book.Id),
		UUID:   book.UUID,
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
	}
	return bookRes, nil
}

func (s *BookService) CreateBook(book *request.CreateBookReq) (bookCreate *response.BookResponse, err error) {
	validator := s.Validate.Struct(book)
	helpers.ErrorPanic(validator)

	bookModel := models.Book{}
	createdBook, err := s.repo.Create(&bookModel)
	if err != nil {
		return bookCreate, err
	}

	bookCreate = &response.BookResponse{
		Id:     int(createdBook.Id),
		UUID:   createdBook.UUID,
		Title:  createdBook.Title,
		Author: createdBook.Author,
		Year:   createdBook.Year,
	}
	return bookCreate, nil
}

func (s *BookService) UpdateBook(bookReq *request.UpdateBookReq) (*response.BookResponse, error) {
	book, err := s.repo.GetByID(bookReq.UUID)
	helpers.ErrorPanic(err)

	// book.Title = request.Title
	// book.Author = request.Author
	// book.Year = request.Year
	s.repo.Update(book)

	bookRes := response.BookResponse{
		Id:     int(book.Id),
		UUID:   book.UUID,
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
	}

	return &bookRes, nil
}

func (s *BookService) DeleteBook(uuid string) error {
	return s.repo.Delete(uuid)
}
