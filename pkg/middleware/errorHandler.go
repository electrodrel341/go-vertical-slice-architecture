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
			if e, ok := err.(*apperror.AppError); ok {
				logger.Error().
					Str("source", "error-handler").
					Str("error_id", e.Id.UUID.String()).
					Str("message", e.Message).
					Int("http_code", e.Code).
					Int("internal_code", e.InternalCode).
					Str("description", e.InternalCodeDescription).
					Bytes("stack", e.StackTrace).
					Err(e.Cause).
					Msg("AppError occurred")

				return c.Status(e.Code).JSON(messages.ErrorResponseAppError(e))
			}

			logger.Error().
				Str("source", "error-handler").
				Err(err).
				Msg("Unexpected error")

			return c.Status(fiber.StatusInternalServerError).JSON(messages.ErrorResponse(err))
		}

		return nil
	}
}
