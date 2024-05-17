package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := ValidateJWT(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			ctx.Abort()
			return
		}

		error := ValidateAdminRoleJWT(ctx)
		if error != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Only Administrator is allowed to perform this action"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func JWTAuthCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := ValidateJWT(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			ctx.Abort()
			return
		}

		error := ValidateCustomerRoleJWT(ctx)
		if error != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Only registered Customers are allowed to perform this action"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
