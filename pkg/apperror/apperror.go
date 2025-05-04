package apperror

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"runtime/debug"
)

type AppError struct {
	Code                    int
	Message                 string
	InternalCode            int
	InternalCodeDescription string
	Id                      uuid.NullUUID
	StackTrace              []byte
	Cause                   error
}

func (r *AppError) Error() string {
	return fmt.Sprintf("AppError[%s]: %s (HTTP %d, InternalCode %d - %s)",
		r.Id.UUID.String(), r.Message, r.Code, r.InternalCode, r.InternalCodeDescription)
}

func (e *AppError) Unwrap() error {
	return e.Cause
}

func NewAppError(HTTPCode int, errorData ErrorData, cause error) *AppError {
	return &AppError{
		Code:                    HTTPCode,
		Message:                 errorData.Message,
		InternalCode:            errorData.CodeValue,
		InternalCodeDescription: errorData.CodeDescription,
		Id:                      uuid.NullUUID{UUID: uuid.New()},
		StackTrace:              debug.Stack(),
		Cause:                   cause,
	}
}

//func NotFound(Message string) *AppError {
//	return NewAppError(http.StatusNotFound, Message)
//}

func BadRequest(errorData ErrorData, cause error) *AppError {
	return NewAppError(http.StatusBadRequest, errorData, cause)
}

//func Unauthorized(Message string) *AppError {
//	return NewAppError(http.StatusUnauthorized, Message)
//}

//func Forbidden(Message string) *AppError {
//	return NewAppError(http.StatusForbidden, Message)
//}

func InternalServerError(cause error) *AppError {
	return NewAppError(http.StatusInternalServerError, ErrorInternalServerError, cause)
}

func EntityNotFound(errorData ErrorData) *AppError {
	return NewAppError(http.StatusUnprocessableEntity, errorData, nil)
}

func ConfigNotFound(errorData ErrorData) *AppError {
	return NewAppError(http.StatusInternalServerError, errorData, nil)
}
