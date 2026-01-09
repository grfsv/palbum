package presentation

import (
	"encoding/json"
	"fmt"
	"net/http"
	"remind_map/internal/domain/commons"

	"remind_map/internal/utils/log"

	"github.com/cockroachdb/errors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

	err = BindErrorConvert(err)

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

// BindErrorConvert : ShouldBindJSONのエラーを map[string]string に変換する
func BindErrorConvert(err error) error {
	var messages []string

	// 1. バリデーションエラーの場合（ルール違反）
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		for _, ve := range validationErrors {
			messages = append(messages, fmt.Sprintf("%s is %s", ve.Field(), ve.Tag()))
		}
	}

	// 2. JSONの型違い（数値フィールドに文字列を入れたなど）
	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		messages = append(messages, fmt.Sprintf("field '%s' has invalid type", unmarshalTypeError.Field))
	}

	if len(messages) == 1 {
		err = commons.NewValidationError(messages[0])
	} else if len(messages) > 1 {
		var multiErr []error
		for _, msg := range messages {
			multiErr = append(multiErr, commons.NewValidationError(msg))
		}
		err = commons.NewMultipleValidationError(multiErr)
	}

	return err
}
