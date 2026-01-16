package presentation

import (
	"errors"
	"net/http"
	usecase "palbum/internal/application/usecase/friend"
	domain "palbum/internal/domain/friend"
	"palbum/internal/presentation/middleware"
	"palbum/internal/presentation/utils"

	"github.com/gin-gonic/gin"
)

type FriendHandler struct {
	friendCode    *usecase.FriendCodeUsecase
	friendRequest *usecase.FriendRequestUsecase
	requestStatus *usecase.RequestStatusUsecase
	friendList    *usecase.FriendListUsecase
	errorHandler  *utils.ErrorHandler
}

func NewFriendHandler(
	friendCode *usecase.FriendCodeUsecase,
	friendRequest *usecase.FriendRequestUsecase,
	requestStatus *usecase.RequestStatusUsecase,
	friendList *usecase.FriendListUsecase,
	errorHandler *utils.ErrorHandler,
) *FriendHandler {
	return &FriendHandler{
		friendCode:    friendCode,
		friendRequest: friendRequest,
		requestStatus: requestStatus,
		friendList:    friendList,
		errorHandler:  errorHandler,
	}
}

func (h *FriendHandler) GetFriendCode(ctx *gin.Context) {
	userUUID := middleware.MustGetUserUUID(ctx)
	usecaseInput := usecase.FriendCodeInput{
		UserUUID: userUUID,
	}

	res, err := h.friendCode.Execute(ctx.Request.Context(), usecaseInput)
	if err != nil {
		h.errorHandler.HandleError(ctx, err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func (h *FriendHandler) RequestFriend(ctx *gin.Context) {
	var uri struct {
		Code string `binding:"required" uri:"friend_code"`
	}

	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		h.errorHandler.HandleError(ctx, err)

		return
	}

	code, err := domain.NewCodeFromString(uri.Code)
	if err != nil {
		h.errorHandler.HandleError(ctx, err)

		return
	}

	userUUID := middleware.MustGetUserUUID(ctx)
	input := usecase.FriendRequestRequest{
		UserUUID: userUUID,
		Code:     code,
	}

	err = h.friendRequest.Execute(ctx.Request.Context(), input)
	if err != nil {
		h.errorHandler.HandleError(ctx, err)

		return
	}

	ctx.Status(http.StatusOK)
}

func (h *FriendHandler) UpdateRequestStatus(ctx *gin.Context) {
	var uri struct {
		RequestUUID string `binding:"required" uri:"request_uuid"`
	}

	var body struct {
		Status string `binding:"required" json:"status"`
	}

	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		h.errorHandler.HandleError(ctx, err)

		return
	}

	err = ctx.ShouldBindJSON(&body)
	if err != nil {
		h.errorHandler.HandleError(ctx, err)

		return
	}

	var validations []error

	status, err := domain.ParseStatus(body.Status)
	if err != nil {
		validations = append(validations, err)
	}

	requestUUID, err := domain.NewUUIDFromString(uri.RequestUUID)
	if err != nil {
		validations = append(validations, err)
	}

	if len(validations) > 0 {
		h.errorHandler.HandleError(ctx, errors.Join(validations...))

		return
	}

	userUUID := middleware.MustGetUserUUID(ctx)
	input := usecase.RequestStatusInput{
		UserUUID:    userUUID,
		RequestUUID: requestUUID,
		Status:      status,
	}

	err = h.requestStatus.Execute(ctx.Request.Context(), input)
	if err != nil {
		h.errorHandler.HandleError(ctx, err)

		return
	}

	ctx.Status(http.StatusOK)
}

func (h *FriendHandler) GetFriendList(ctx *gin.Context) {
	userUUID := middleware.MustGetUserUUID(ctx)

	input := usecase.FriendListInput{
		UserUUID: userUUID,
	}

	res, err := h.friendList.Execute(ctx.Request.Context(), input)
	if err != nil {
		h.errorHandler.HandleError(ctx, err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
