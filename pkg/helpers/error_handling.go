package helpers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/awahids/belajar-go/internal/delivery/data/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrorValidator(err error) error {
	if err != nil {
		return fmt.Errorf("validation error: %v", err)
	}
	return nil
}

func JSONResponse(ctx *gin.Context, data interface{}) {
	isCreate := ctx.Request.Method == http.MethodPost
	statusCode := http.StatusOK
	if isCreate {
		statusCode = http.StatusCreated
	}

	ctx.JSON(statusCode, data)
}

var ErrRecordNotFound = gorm.ErrRecordNotFound

func ErrorResponse(err error, ctx *gin.Context) {
	// duplicate key error
	if strings.Contains(err.Error(), "duplicate key") {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
			Data:    "User already exists",
		}
		JSONResponse(ctx, webResponse)
		return
	}

	// default error
	if err != nil {
		var webResponse response.Response
		if errors.Is(err, ErrRecordNotFound) {
			webResponse = response.Response{
				Code:    http.StatusNotFound,
				Message: "Not Found",
				Data:    "Data not found",
			}
		} else {
			webResponse = response.Response{
				Code:    http.StatusInternalServerError,
				Message: "Internal Server Error",
				Data:    ErrRecordNotFound.Error(),
			}
		}
		JSONResponse(ctx, webResponse)
		return
	}
}
