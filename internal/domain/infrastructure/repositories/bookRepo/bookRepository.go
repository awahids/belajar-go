package bookRepo

import (
	"errors"

	"github.com/awahids/belajar-gin/internal/domain/infrastructure/repositories/repoInterface"
	"github.com/awahids/belajar-gin/internal/domain/models"
	"github.com/awahids/belajar-gin/pkg/helpers"
	"gorm.io/gorm"
)

type BookRepository struct {
	Db *gorm.DB
}

func NewBookRepository(Db *gorm.DB) repoInterface.BookInterface {
	return &BookRepository{Db: Db}
}

func (r *BookRepository) GetAll() []models.Book {
	var books []models.Book
	results := r.Db.Find(&books)
	helpers.ErrorPanic(results.Error)
	return books
}

func (r *BookRepository) GetByID(uuid string) (books models.Book, err error) {
	var book models.Book
	result := r.Db.Find(&book, uuid)

	if result != nil {
		return book, nil
	} else {
		return book, errors.New("book is not found")
	}
}

func (r *BookRepository) Create(book models.Book) {
	results := r.Db.Create(&book)
	helpers.ErrorPanic(results.Error)
}

func (r *BookRepository) Update(book models.Book) {
	results := r.Db.Updates(&book)
	helpers.ErrorPanic(results.Error)
}

func (r *BookRepository) Delete(uuid string) error {
	results := r.Db.Where("uuid = ?", uuid).Delete(&models.Book{})
	if results.Error != nil {
		return results.Error
	}

	return nil
}
