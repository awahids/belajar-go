package repoInterface

import (
	"github.com/awahids/belajar-gin/internal/domain/models"
)

type BookInterface interface {
	Create(book models.Book)
	Update(book models.Book)
	GetByID(uuid string) (book models.Book, err error)
	GetAll() []models.Book
	Delete(uuid string) error
}
