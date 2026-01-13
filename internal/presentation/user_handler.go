package presentation

import (
	"net/http"
	usecase "palbum/internal/application/usecase/user"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	signup       *usecase.SignUpUsecase
	login        *usecase.UserLoginUsecase
	logout       *usecase.UserLogoutUsecase
	refresh      *usecase.RefreshUsecase
	errorHandler *ErrorHandler
}

func NewUserHandler(
	signup *usecase.SignUpUsecase,
	login *usecase.UserLoginUsecase,
	logout *usecase.UserLogoutUsecase,
	refresh *usecase.RefreshUsecase,
	errorHandler *ErrorHandler,
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
	var req usecase.SignUpRequest

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
	var req usecase.UserLoginRequest

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
	var req usecase.UserLogoutRequest

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
	var req usecase.RefreshRequest

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
