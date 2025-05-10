package injection

import (
	llmhandler "PetAi/internal/llm/handler"
	producthandler "PetAi/internal/product/handler"
	userhandler "PetAi/internal/user/handler"
	"PetAi/pkg/config"
	"PetAi/pkg/middleware"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func InvokeApp() error {
	return container.Invoke(func(
		app *fiber.App,
		jwtMw *middleware.Middleware,
	) {
		// create health end point
		app.Get("/health", func(c *fiber.Ctx) error {
			return c.SendString("Status ok - api running")
		})

		api := app.Group("/api")

		userApi := api.Group("/user")
		userhandler.UserRouter(userApi, UserServiceProvider)

		productApi := api.Group("/product")
		producthandler.ProductRouter(productApi, ProductServiceProvider)

		llmApi := api.Group("/llm")
		llmhandler.LLMRouter(llmApi, LLMServiceProvider, jwtMw)

		err := app.Listen(fmt.Sprintf(":%s", config.Get().APPConfig.Port))
		if err != nil {
			log.Fatal().Err(err).Msg("Ошибка запуска сервиса")
		}
	})
}
