package repoInterface

import (
	"github.com/awahids/belajar-gin/internal/domain/models"
)

type BookInterface interface {
	Create(book models.Book) (bookCreate models.Book, err error)
	Update(book models.Book) (updatedBook models.Book)
	GetByID(uuid string) (book models.Book, err error)
	GetAll() (books []models.Book, err error)
	Delete(uuid string) error
}
