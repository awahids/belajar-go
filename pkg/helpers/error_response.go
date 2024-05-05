package helpers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/awahids/belajar-go/internal/delivery/data/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var ErrRecordNotFound = gorm.ErrRecordNotFound

func ErrorResponse(err error, ctx *gin.Context) {
	// duplicate key error
	if strings.Contains(err.Error(), "duplicate key") {
		webResponse := response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   "User already exists",
		}
		JSONResponse(ctx, webResponse)
		return
	}

	// default error
	if err != nil {
		var webResponse response.Response
		if errors.Is(err, ErrRecordNotFound) {
			webResponse = response.Response{
				Code:   http.StatusNotFound,
				Status: "Not Found",
				Data:   "Data not found",
			}
		} else {
			webResponse = response.Response{
				Code:   http.StatusInternalServerError,
				Status: "Internal Server Error",
				Data:   ErrRecordNotFound.Error(),
			}
		}
		JSONResponse(ctx, webResponse)
		return
	}
}
