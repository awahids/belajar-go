package helpers

import (
	"github.com/go-playground/validator/v10"
)

func ErrorValidator(err error) error {
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			return err // Jika bukan error validasi, kembalikan seperti biasa
		}
		for _, valErr := range validationErrors {
			return valErr // Mengembalikan hanya error pertama
		}
	}
	return nil
}
