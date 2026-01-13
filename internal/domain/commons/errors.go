package commons

type ErrorType int

const (
	TypeNotFound ErrorType = iota
	TypeValidation
	TypeConflict
	TypeUnAuthorized
)

func (t ErrorType) String() string {
	switch t {
	case TypeNotFound:
		return "NotFound"
	case TypeValidation:
		return "Validation"
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
	return &CustomError{
		Type:    TypeNotFound,
		Message: message,
		Cause:   nil,
		Object:  nil,
	}
}

func NewValidationError(message string) error {
	return &CustomError{
		Type:    TypeValidation,
		Message: message,
		Cause:   nil,
		Object:  nil,
	}
}

func NewConflictError(message string) error {
	return &CustomError{
		Type:    TypeConflict,
		Message: message,
		Cause:   nil,
		Object:  nil,
	}
}

func NewUnAuthorizedError() error {
	return &CustomError{
		Type:    TypeUnAuthorized,
		Message: "unauthorized",
		Cause:   nil,
		Object:  nil,
	}
}

func (e *CustomError) Error() string {
	if e.Cause != nil {
		return e.Message + ": " + e.Cause.Error()
	}

	return e.Message
}
