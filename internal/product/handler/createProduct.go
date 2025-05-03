package handler

import (
	"PetAi/internal/product"
	"PetAi/internal/product/service"
	"PetAi/pkg/apperror"
	"PetAi/pkg/messages"
	"PetAi/pkg/validate"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// Body request schema for CreateProduct
type ProductSchema struct {
	Name     string  `json:"name" validate:"required,min=5"`
	Sku      string  `json:"sku" validate:"required,min=8"`
	Category string  `json:"category" validate:"required,min=5"`
	Price    float64 `json:"price" validate:"required"`
}

// Creates a new product into the database
func CreateProduct(s *service.CreateProductService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get body request
		var body ProductSchema
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

		// No schema errores then map body to domain
		p := &product.Product{
			Name:     body.Name,
			Sku:      body.Sku,
			Category: product.ParseProductCategory(body.Category),
			Price:    body.Price,
		}

		// Execute the service
		result, err := s.CreateProduct(p)
		if err != nil {
			// if service response an error return via the middleware
			log.Error(err)
			return err
		}

		// Success execution
		return c.Status(fiber.StatusCreated).JSON(messages.SuccessResponse(&result))
	}
}
