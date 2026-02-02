package presentation

import (
	"net/http"
	"time"

	usecase "palbum/internal/application/usecase/dtk"
	"palbum/internal/presentation/middleware"
	"palbum/internal/presentation/utils"

	"github.com/gin-gonic/gin"
)

type DTKHandler struct {
	getMyDTK      *usecase.GetMyDTKUsecase
	getFriendDTKs *usecase.GetFriendDTKsUsecase
	errorHandler  *utils.ErrorHandler
}

func NewDTKHandler(
	getMyDTK *usecase.GetMyDTKUsecase,
	getFriendDTKs *usecase.GetFriendDTKsUsecase,
	errorHandler *utils.ErrorHandler,
) *DTKHandler {
	return &DTKHandler{
		getMyDTK:      getMyDTK,
		getFriendDTKs: getFriendDTKs,
		errorHandler:  errorHandler,
	}
}

func (h *DTKHandler) GetMyDTK(ctx *gin.Context) {
	userUUID := middleware.MustGetUserUUID(ctx)

	input := usecase.GetMyDTKInput{
		UserUUID: userUUID,
	}

	res, err := h.getMyDTK.Execute(ctx.Request.Context(), input)
	if err != nil {
		h.errorHandler.HandleError(ctx, err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func (h *DTKHandler) GetFriendDTKs(ctx *gin.Context) {
	userUUID := middleware.MustGetUserUUID(ctx)

	var since *time.Time

	sinceStr := ctx.Query("since")
	if sinceStr != "" {
		reqDate, err := time.Parse("2006-01-02", sinceStr)
		if err != nil {
			h.errorHandler.HandleError(ctx, err)

			return
		}

		since = &reqDate
	}

	input := usecase.GetFriendDTKsInput{
		UserUUID: userUUID,
		Since:    since,
	}

	res, err := h.getFriendDTKs.Execute(ctx.Request.Context(), input)
	if err != nil {
		h.errorHandler.HandleError(ctx, err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
