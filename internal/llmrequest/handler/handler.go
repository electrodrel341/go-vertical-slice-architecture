package handler

import (
	"PetAi/internal/llmrequest/service"
	"github.com/gofiber/fiber/v2"
)

// LLMRouter is the Router for GoFiber App
func LLMRouter(app fiber.Router, s *service.LLMService) {
	app.Post("/", SendRequest(s.SendLLMRequestServiceProvider))
}
