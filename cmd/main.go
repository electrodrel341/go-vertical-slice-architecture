package main

import (
	llmhandler "PetAi/internal/llmrequest/handler"
	producthandler "PetAi/internal/product/handler"
	userhandler "PetAi/internal/user/handler"
	"PetAi/pkg/config"
	"PetAi/pkg/injection"
	"PetAi/pkg/middleware"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Ошибка загрузки .env файла")
	}

	err = config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// prepare all components for dependency injection
	injection.ProvideComponents()

	// initiate service components with dependency injection
	if err := injection.InitComponents(); err != nil {
		panic(err)
	}

	// create fiber
	app := fiber.New()

	// add fiber middlewares
	app.Use(cors.New())
	app.Use(helmet.New())

	// custom middlewares
	app.Use(middleware.ErrorHandler)

	// create health end point
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Status ok - api running")
	})

	// create api group
	api := app.Group("/api") // /api

	// add api group for user
	userApi := api.Group("/user") // /api/user
	userhandler.UserRouter(userApi, injection.UserServiceProvider)

	// add api group for product
	productApi := api.Group("/product") // /api/product
	producthandler.ProductRouter(productApi, injection.ProductServiceProvider)

	// add api group for product
	llmApi := api.Group("/llm") // /api/product
	llmhandler.LLMRouter(llmApi, injection.LLMServiceProvider)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Get().APPConfig.Port)))
}
