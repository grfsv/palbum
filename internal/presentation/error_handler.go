package presentation

import (
	"fmt"
	"net/http"
	"palbum/internal/domain/commons"

	"palbum/internal/utils/log"

	"github.com/cockroachdb/errors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

type ErrorHandler struct {
	logger *log.Log
}

func NewErrorHandler(logger *log.Log) *ErrorHandler {
	return &ErrorHandler{
		logger: logger,
	}
}

func (h *ErrorHandler) HandleError(ctx *gin.Context, err error) {
	h.logger.Error(
		"An error occurred:",
		"request_id", requestid.Get(ctx),
		"error", err.Error(), // 短いエラーメッセージ
		"stack_trace", fmt.Sprintf("%+v", err),
	)

	var customErr *commons.CustomError
	if errors.As(err, &customErr) {
		switch customErr.Type {
		case commons.TypeValidation:
			ctx.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": customErr.Error(),
				},
			)
		case commons.TypeNotFound:
			ctx.JSON(
				http.StatusNotFound,
				gin.H{
					"error": customErr.Error(),
				})
		case commons.TypeConflict:
			ctx.JSON(
				http.StatusConflict,
				gin.H{
					"error": customErr.Error(),
				})
		case commons.TypeUnAuthorized:
			ctx.JSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "unauthorized",
				})
		}

		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
}
