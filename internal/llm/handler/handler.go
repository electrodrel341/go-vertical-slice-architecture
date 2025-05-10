package handler

import (
	"PetAi/internal/llm/service"
	"PetAi/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

// LLMRouter is the Router for GoFiber App
func LLMRouter(app fiber.Router, s *service.LLMService, jwt *middleware.Middleware) {
	app.Post("/", jwt.AuthRequired(), SendRequest(s.SendLLMRequestServiceProvider))
}
