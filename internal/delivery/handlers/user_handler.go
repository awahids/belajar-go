package handlers

import (
	"net/http"

	"github.com/awahids/belajar-gin/internal/delivery/data/request"
	"github.com/awahids/belajar-gin/internal/delivery/data/response"
	"github.com/awahids/belajar-gin/internal/domain/services/user_service"
	"github.com/awahids/belajar-gin/pkg/helpers"
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
	createUserReq := request.CreateUserReq{}
	err := ctx.ShouldBindJSON(&createUserReq)
	helpers.ErrorPanic(err)

	createdUser, err := h.userService.CreateUser(&createUserReq)
	if err != nil {
		helpers.ErrorResponse(err, ctx)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusCreated,
		Status: "Ok",
		Data:   createdUser,
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
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userRes,
	}
	helpers.JSONResponse(ctx, webResponse)
}

func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	users, err := h.userService.GetAllUsers()
	helpers.ErrorPanic(err)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   users,
	}
	helpers.JSONResponse(ctx, webResponse)
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	updateUserReq := request.UpdateUserReq{}
	err := ctx.ShouldBindJSON(&updateUserReq)
	helpers.ErrorPanic(err)

	userUUID := ctx.Param("uuid")
	updateUserReq.UUID = userUUID

	updatedUser, err := h.userService.UpdateUser(updateUserReq)
	helpers.ErrorPanic(err)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   updatedUser,
	}
	helpers.JSONResponse(ctx, webResponse)
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	userUUID := ctx.Param("uuid")

	err := h.userService.DeleteUser(userUUID)
	helpers.ErrorPanic(err)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	helpers.JSONResponse(ctx, webResponse)
}
