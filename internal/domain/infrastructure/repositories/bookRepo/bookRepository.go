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

func (r *BookRepository) GetAll() ([]models.Book, error) {
	var books []models.Book
	results := r.Db.Find(&books)
	helpers.ErrorPanic(results.Error)
	return books, nil
}

func (r *BookRepository) GetByID(uuid string) (book models.Book, err error) {
	result := r.Db.Where("uuid = ?", uuid).First(&book)
	helpers.ErrorPanic(result.Error)

	if result.Error != nil {
		return book, result.Error
	} else if result.RowsAffected == 0 {
		return book, errors.New("book not found")
	}

	return book, nil
}

func (r *BookRepository) Create(book models.Book) (models.Book, error) {
	result := r.Db.Create(&book)
	if result.Error != nil {
		helpers.ErrorPanic(result.Error)
		return models.Book{}, result.Error
	}
	return book, nil
}

func (r *BookRepository) Update(book models.Book) models.Book {
	results := r.Db.Updates(&book)
	helpers.ErrorPanic(results.Error)

	return book
}

func (r *BookRepository) Delete(uuid string) error {
	results := r.Db.Where("uuid = ?", uuid).Delete(&models.Book{})
	if results.Error != nil {
		return results.Error
	}

	return nil
}
