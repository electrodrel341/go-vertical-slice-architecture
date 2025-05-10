package apperror

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"runtime/debug"
)

type AppErrorInterface interface {
	error
	Unwrap() error
	LogFields() map[string]interface{}
}

type AppError struct {
	Code                    int
	Message                 string
	InternalCode            int
	InternalCodeDescription string
	Id                      uuid.UUID
	StackTrace              []byte
	Cause                   error
}

func (r *AppError) Error() string {
	return fmt.Sprintf("AppError[%s]: %s (HTTP %d, InternalCode %d - %s)",
		r.Id.String(), r.Message, r.Code, r.InternalCode, r.InternalCodeDescription)
}

func (e *AppError) Unwrap() error {
	return e.Cause
}

func (e *AppError) LogFields() map[string]interface{} {
	return map[string]interface{}{
		"error_id":      e.Id.String(),
		"message":       e.Message,
		"http_code":     e.Code,
		"internal_code": e.InternalCode,
		"description":   e.InternalCodeDescription,
		"stack":         string(e.StackTrace),
		"cause":         e.Cause,
	}
}

func NewAppError(HTTPCode int, errorData ErrorData, cause error) *AppError {
	return &AppError{
		Code:                    HTTPCode,
		Message:                 errorData.Message,
		InternalCode:            errorData.CodeValue,
		InternalCodeDescription: errorData.CodeDescription,
		Id:                      uuid.New(),
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

func Unauthorized(cause error) *AppError {
	return NewAppError(http.StatusUnauthorized, ErrorUnauthorized, cause)
}

func LoginError(errorData ErrorData) *AppError {
	return NewAppError(http.StatusUnauthorized, errorData, nil)
}

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

func EntityDuplicate(errorData ErrorData) *AppError {
	return NewAppError(http.StatusConflict, errorData, nil)
}
