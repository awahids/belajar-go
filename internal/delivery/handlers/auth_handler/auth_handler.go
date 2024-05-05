package auth_handler

import (
	"net/http"

	"github.com/awahids/belajar-gin/internal/delivery/data/request"
	"github.com/awahids/belajar-gin/internal/domain/services/auth_service"
	"github.com/awahids/belajar-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *auth_service.AuthService
}

func NewAuthHandler(authService *auth_service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var loginReq request.LoginReq
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authService.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Send token as response
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Login successful",
		"token":      token,
	})
}
