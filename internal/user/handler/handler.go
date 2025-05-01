package handler

import (
	"PetAi/internal/user/service"
	"github.com/gofiber/fiber/v2"
)

// UserRouter is the Router for GoFiber App
func UserRouter(app fiber.Router, s *service.UserService) {
	app.Post("/", CreateUser(s.CreateUserServiceProvider))
}
