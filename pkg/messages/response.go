package messages

import (
	"PetAi/pkg/apperror"
	"github.com/gofiber/fiber/v2"
)

type ResponseErrorData struct {
	Message                 string `json:"message"`
	InternalCode            int    `json:"internal_code"`
	InternalCodeDescription string `json:"internal_code_description"`
	Id                      string `json:"id"`
}

// SuccessResponseSlice is the list SuccessResponse that will be passed in the response by Handler
func SuccessResponseSlice(data *[]any) *fiber.Map {
	return &fiber.Map{
		"status":     true,
		"data":       data,
		"error":      nil,
		"error_data": nil,
	}
}

// SuccessResponse is the primitive type SuccessResponse that will be passed in the response by Handler
func SuccessResponse(data any) *fiber.Map {
	return &fiber.Map{
		"status":     true,
		"data":       data,
		"error":      nil,
		"error_data": nil,
	}
}

// ErrorResponse is the ErrorResponse that will be passed in the response by Handler
func ErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status":     false,
		"data":       nil,
		"error":      err.Error(),
		"error_data": nil,
	}
}

func ErrorResponseAppError(err *apperror.AppError) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   nil,
		"error":  err.Error(),
		"error_data": ResponseErrorData{
			Id:                      err.Id.String(),
			Message:                 err.Message,
			InternalCode:            err.InternalCode,
			InternalCodeDescription: err.InternalCodeDescription,
		},
	}
}
