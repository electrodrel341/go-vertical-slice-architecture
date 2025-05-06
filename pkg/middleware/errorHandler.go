package middleware

import (
	"PetAi/pkg/apperror"
	"PetAi/pkg/messages"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

// ErrorHandler is a middleware that converts AppError to fiber.Error.

func ErrorHandler(logger zerolog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()

		if err != nil {
			apperror.LogError(logger, "error-handler", err)

			if e, ok := err.(*apperror.AppError); ok {
				return c.Status(e.Code).JSON(messages.ErrorResponseAppError(e))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(messages.ErrorResponse(err))
		}

		return nil
	}
}
