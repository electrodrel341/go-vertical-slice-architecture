package handler

import (
	"PetAi/internal/llm"
	"PetAi/internal/llm/service"
	"PetAi/pkg/apperror"
	"PetAi/pkg/messages"
	"PetAi/pkg/validate"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type LLMRequestSchema struct {
	Message    string `json:"message" validate:"required,min=5"`
	AIProvider string `json:"ai_provider" validate:"required,min=2"`
}

// Send a new request into the LLM
func SendRequest(s *service.SendLLMRequestService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get body request
		var body LLMRequestSchema
		// Validate the body
		err := c.BodyParser(&body)
		if err != nil {
			// Map the error and response via the middleware
			log.Error(err)
			return err
		}

		// Validate schema
		appErr, err := validate.Validate(body)
		if err != nil {
			log.Error(appErr)
			return apperror.BadRequest(appErr, err)
		}

		aiApiProvider, err := llm.ParseAIApiProvider(body.AIProvider)
		if err != nil {
			// if service response an error return via the middleware
			log.Error(err)
			return err
		}

		// No schema errores then map body to domain
		model := llm.O4mini
		p := llm.NewLLMRequest(
			aiApiProvider,
			llm.NewPromt(body.Message, &model, nil, nil, nil, "You are a helpful assistant that creates blog outlines.", nil, nil),
		)

		// Execute the service
		result, err := s.SendRequest(p)
		if err != nil {
			// if service response an error return via the middleware
			log.Error(err)
			return err
		}

		// Success execution
		return c.Status(fiber.StatusOK).JSON(messages.SuccessResponse(&result))
	}
}
