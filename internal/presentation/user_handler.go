package presentation

import (
	"net/http"
	"palbum/internal/application/usecase/user"
	"palbum/internal/presentation/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	signup       *user.SignUpUsecase
	login        *user.UserLoginUsecase
	logout       *user.UserLogoutUsecase
	refresh      *user.RefreshUsecase
	errorHandler *utils.ErrorHandler
}

func NewUserHandler(
	signup *user.SignUpUsecase,
	login *user.UserLoginUsecase,
	logout *user.UserLogoutUsecase,
	refresh *user.RefreshUsecase,
	errorHandler *utils.ErrorHandler,
) *UserHandler {
	return &UserHandler{
		signup:       signup,
		login:        login,
		logout:       logout,
		refresh:      refresh,
		errorHandler: errorHandler,
	}
}

func (h *UserHandler) SignUp(ctx *gin.Context) {
	var req user.SignUpRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		h.errorHandler.HandleError(ctx, err)

		return
	}

	res, err := h.signup.Execute(ctx.Request.Context(), req)
	if err != nil {
		h.errorHandler.HandleError(ctx, err)

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": res,
	})
}

func (h *UserHandler) Login(ctx *gin.Context) {
	var req user.UserLoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		h.errorHandler.HandleError(ctx, err)

		return
	}

	res, err := h.login.Execute(ctx.Request.Context(), req)
	if err != nil {
		h.errorHandler.HandleError(ctx, err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func (h *UserHandler) Logout(ctx *gin.Context) {
	var req user.UserLogoutRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		h.errorHandler.HandleError(ctx, err)

		return
	}

	err = h.logout.Execute(ctx.Request.Context(), req)
	if err != nil {
		h.errorHandler.HandleError(ctx, err)

		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *UserHandler) Refresh(ctx *gin.Context) {
	var req user.RefreshRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		h.errorHandler.HandleError(ctx, err)

		return
	}

	res, err := h.refresh.Execute(ctx.Request.Context(), req)
	if err != nil {
		h.errorHandler.HandleError(ctx, err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
