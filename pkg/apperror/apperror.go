package apperror

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type AppError struct {
	Code                    int
	Message                 string
	InternalCode            int
	InternalCodeDescription string
	Id                      uuid.NullUUID
}

func (r *AppError) Error() string {
	return fmt.Sprintf(r.Message)
}

func NewAppError(HTTPCode int, errorData ErrorData) *AppError {
	return &AppError{
		Code:                    HTTPCode,
		Message:                 errorData.Message,
		InternalCode:            errorData.CodeValue,
		InternalCodeDescription: errorData.CodeDescription,
		Id:                      uuid.NullUUID{UUID: uuid.New()},
	}
}

//func NotFound(Message string) *AppError {
//	return NewAppError(http.StatusNotFound, Message)
//}

func BadRequest(errorData ErrorData) *AppError {
	return NewAppError(http.StatusBadRequest, errorData)
}

//func Unauthorized(Message string) *AppError {
//	return NewAppError(http.StatusUnauthorized, Message)
//}

//func Forbidden(Message string) *AppError {
//	return NewAppError(http.StatusForbidden, Message)
//}

func InternalServerError() *AppError {
	return NewAppError(http.StatusInternalServerError, ErrorInternalServerError)
}

func EntityNotFound(errorData ErrorData) *AppError {
	return NewAppError(http.StatusUnprocessableEntity, errorData)
}

func ConfigNotFound(errorData ErrorData) *AppError {
	return NewAppError(http.StatusInternalServerError, errorData)
}
