package commons

import (
	"github.com/cockroachdb/errors"
)

type ErrorType int

const (
	TypeNotFound ErrorType = iota
	TypeValidation
	TypeMultipleValidation
	TypeConflict
	TypeUnAuthorized
)

func (t ErrorType) String() string {
	switch t {
	case TypeNotFound:
		return "NotFound"
	case TypeValidation:
		return "Validation"
	case TypeMultipleValidation:
		return "MultipleValidation"
	case TypeConflict:
		return "Conflict"
	case TypeUnAuthorized:
		return "UnAuthorized"
	default:
		return "Unknown"
	}
}

type CustomError struct {
	Type    ErrorType
	Message string
	Cause   error
	Object  any
}

func NewNotFoundError(message string) error {
	return errors.WithStack(&CustomError{
		Type:    TypeNotFound,
		Message: message,
		Cause:   nil,
		Object:  nil,
	})
}

func NewValidationError(message string) error {
	return errors.WithStack(&CustomError{
		Type:    TypeValidation,
		Message: message,
		Cause:   nil,
		Object:  nil,
	})
}

func NewConflictError(message string) error {
	return errors.WithStack(&CustomError{
		Type:    TypeConflict,
		Message: message,
		Cause:   nil,
		Object:  nil,
	})
}

func NewUnAuthorizedError() error {
	return errors.WithStack(&CustomError{
		Type:    TypeUnAuthorized,
		Message: "unauthorized",
		Cause:   nil,
		Object:  nil,
	})
}

func NewMultipleValidationError(errorList []error) error {
	if len(errorList) <= 1 {
		return errors.New("NewMultipleValidationError: no enough length errors")
	}

	var customErr *CustomError

	for _, err := range errorList {
		if errors.As(err, &customErr) {
			if customErr.Type != TypeValidation {
				return errors.New("NewMultipleValidationError: include invalid error type")
			}
		} else {
			return errors.New("NewMultipleValidationError: include non CustomError type")
		}
	}

	return errors.WithStack(
		&CustomError{
			Type:    TypeMultipleValidation,
			Message: "multiple validation errors",
			Cause:   nil,
			Object:  errorList,
		})
}

func (e *CustomError) Error() string {
	if e.Cause != nil {
		return e.Message + ": " + e.Cause.Error()
	}

	return e.Message
}

// func (e *CustomError) LogValue() slog.Value {
// 	// 基本属性
// 	attrs := []slog.Attr{
// 		slog.String("type", e.Type.String()), // 文字列化したType
// 		slog.String("message", e.Message),
// 	}

// 	// 内部エラー (Cause) がある場合はスタックトレース付きで出力
// 	if e.Cause != nil {
// 		attrs = append(attrs, slog.String("cause", fmt.Sprintf("%+v", e.Cause)))
// 	}

// 	// Object (詳細情報や複数のバリデーションエラー) がある場合
// 	if e.Object != nil {
// 		// slog.Any は、中身が LogValuer を実装していれば
// 		// 再帰的にその LogValue を呼んでくれます。
// 		// つまり、MultipleValidation の場合、子供のエラーもきれいに構造化されます。
// 		attrs = append(attrs, slog.Any("details", e.Object))
// 	}

// 	// グループ化して返す
// 	return slog.GroupValue(attrs...)
// }
