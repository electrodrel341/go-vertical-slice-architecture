package handler

import (
	"PetAi/internal/product/service"
	"github.com/gofiber/fiber/v2"
)

// ProductRouter is the Router for GoFiber App
func ProductRouter(app fiber.Router, s *service.ProductService) {
	app.Post("/", CreateProduct(s.CreateProductServiceProvider))
}
