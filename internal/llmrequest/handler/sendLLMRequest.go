package handler

import (
	"PetAi/internal/llmrequest"
	"PetAi/internal/llmrequest/service"
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
		serr, err := validate.Validate(body)
		if err != nil {
			log.Error(serr)
			return apperror.BadRequest(serr)
		}

		aiApiProvider, err := llmrequest.ParseAIApiProvider(body.AIProvider)
		if err != nil {
			// if service response an error return via the middleware
			log.Error(err)
			return err
		}

		// No schema errores then map body to domain
		p := &llmrequest.LLMRequest{
			AIApiProvider: aiApiProvider,
			Promt: llmrequest.Promt{
				RequestMessage: body.Message,
				SystemMessage:  "You are a helpful assistant that creates blog outlines.",
				Model:          llmrequest.O4mini,
			},
		}

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
