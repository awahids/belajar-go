package middlewares

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/awahids/belajar-go/internal/domain/models"
	"github.com/awahids/belajar-go/internal/domain/repositories/user_repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var privateKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(user *models.User) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// "id":      user.ID,
		"uuid":    user.UUID,
		"role_id": user.RoleId,
		"iat":     time.Now().Unix(),
		"eat":     time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	return token.SignedString(privateKey)
}

func CurrentUser(ctx *gin.Context, db *gorm.DB) *models.User {
	userRepo := user_repository.NewUserRepository(db)

	err := ValidateJWT(ctx)
	if err != nil {
		return &models.User{}
	}
	token, _ := getToken(ctx)
	claims, _ := token.Claims.(jwt.MapClaims)
	userUuuid := claims["uuid"].(string)

	user, err := userRepo.GetUser(userUuuid)
	if err != nil {
		return &models.User{}
	}
	return user
}

func ValidateJWT(ctx *gin.Context) error {
	token, err := getToken(ctx)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")
}

func ValidateAdminRoleJWT(ctx *gin.Context) error {
	token, err := getToken(ctx)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := uint(claims["role_id"].(float64))
	if ok && token.Valid && userRole == 1 {
		return nil
	}
	return errors.New("invalid admin token provided")
}

func ValidateCustomerRoleJWT(ctx *gin.Context) error {
	token, err := getToken(ctx)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := uint(claims["role"].(float64))
	if ok && token.Valid && userRole == 2 || userRole == 1 {
		return nil
	}
	return errors.New("invalid author token provided")
}

func getToken(ctx *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(ctx)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

func getTokenFromRequest(ctx *gin.Context) string {
	bearerToken := ctx.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
