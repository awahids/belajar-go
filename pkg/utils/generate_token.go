package utils

import (
	"os"

	"github.com/awahids/belajar-go/internal/domain/models"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"user_id": int(user.ID),
		"email":   user.Email,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
