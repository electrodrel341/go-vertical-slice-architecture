package injection

import (
	"PetAi/pkg/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/rs/zerolog"
)

func provideFiberApp(logger zerolog.Logger) *fiber.App {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(middleware.RequestLogger(logger))
	app.Use(middleware.ErrorHandler(logger))

	return app
}
