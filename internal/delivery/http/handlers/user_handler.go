package handlers

import (
	"net/http"

	"github.com/awahids/belajar-go/internal/delivery/data/dtos"
	"github.com/awahids/belajar-go/internal/delivery/data/response"
	"github.com/awahids/belajar-go/internal/domain/services/user_service"
	"github.com/awahids/belajar-go/pkg/helpers"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService user_service.UserInterface
}

func NewUserHandler(userInterface user_service.UserInterface) *UserHandler {
	return &UserHandler{
		userService: userInterface,
	}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	createUserReq := dtos.CreateUserReq{}
	err := ctx.ShouldBindJSON(&createUserReq)
	helpers.ErrorPanic(err)

	createdUser, err := h.userService.CreateUser(&createUserReq)
	if err != nil {
		helpers.ErrorResponse(err, ctx)
		return
	}

	webResponse := response.Response{
		Code:    http.StatusCreated,
		Message: "Success Created",
		Data:    createdUser,
	}
	helpers.JSONResponse(ctx, webResponse)
}

func (h *UserHandler) GetUser(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	userRes, err := h.userService.GetUserById(uuid)
	if err != nil {
		helpers.ErrorResponse(err, ctx)
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Message: "Data Found",
		Data:    userRes,
	}
	helpers.JSONResponse(ctx, webResponse)
}

func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	users, err := h.userService.GetAllUsers()
	helpers.ErrorPanic(err)

	webResponse := response.Response{
		Code:    http.StatusOK,
		Message: "Data Found",
		Data:    users,
	}
	helpers.JSONResponse(ctx, webResponse)
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	updateUserReq := dtos.UpdateUserReq{}
	err := ctx.ShouldBindJSON(&updateUserReq)
	helpers.ErrorPanic(err)

	userUUID := ctx.Param("uuid")
	updateUserReq.UUID = userUUID

	updatedUser, err := h.userService.UpdateUser(&updateUserReq)
	helpers.ErrorPanic(err)

	webResponse := response.Response{
		Code:    http.StatusOK,
		Message: "Success Updated",
		Data:    updatedUser,
	}
	helpers.JSONResponse(ctx, webResponse)
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	userUUID := ctx.Param("uuid")

	err := h.userService.DeleteUser(userUUID)
	helpers.ErrorPanic(err)

	webResponse := response.Response{
		Code:    http.StatusOK,
		Message: "Success Deleted",
		Data:    nil,
	}
	helpers.JSONResponse(ctx, webResponse)
}
