package presentation

import (
	"net/http"
	"remind_map/internal/domain/commons"

	"remind_map/internal/utils/log"

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
		"err", err,
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
		case commons.TypeMultipleValidation:
			var messages []string

			errs, ok := customErr.Object.([]error)
			if ok {
				for _, e := range errs {
					messages = append(messages, e.Error())
				}
			}

			ctx.JSON(
				http.StatusBadRequest,
				gin.H{
					"errors": messages,
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
