package injection

import (
	llmservice "PetAi/internal/llm/service"
	productservice "PetAi/internal/product/service"
	userservice "PetAi/internal/user/service"
	"PetAi/pkg/config"
	"PetAi/pkg/database"
	"PetAi/pkg/logger"
	"github.com/rs/zerolog"
	"go.uber.org/dig"
)

// container instance
var container *dig.Container

// new user service instace
var UserServiceProvider *userservice.UserService

// new product service instace
var ProductServiceProvider *productservice.ProductService

var LLMServiceProvider *llmservice.LLMService

// provide components for injection
func ProvideComponents() {
	// create a new container
	container = dig.New()

	err := container.Provide(logger.InitLogger)
	if err != nil {
		panic(err)
	}

	err = container.Provide(provideFiberApp)
	if err != nil {
		panic(err)
	}

	// generate db config instance
	dbConfig := config.Get().DBConfig

	err = container.Provide(func() *database.DbConfig {
		config := database.NewDbConfig(
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Name,
			dbConfig.Host,
			dbConfig.Port,
		)
		config.WithMigration(dbConfig.MigrationsPath)

		return config
	})
	if err != nil {
		panic(err)
	}

	// provide the database connection injection
	err = container.Provide(func(cfg *database.DbConfig, logger zerolog.Logger) *database.DbConn {
		return database.InitPool(cfg, logger)
	})
	if err != nil {
		panic(err)
	}

	// user provider injection
	userservice.ProvideUserComponents(container)

	// product provider injection
	productservice.ProvideProductComponents(container)

	llmservice.ProvideLLMComponents(container)
}

// init service container
func InitComponents() error {
	// user service init componets injection
	UserServiceProvider = userservice.NewUserService()
	err := UserServiceProvider.InitUserComponents(container)
	if err != nil {
		panic(err)
	}

	// product service init componets injection
	ProductServiceProvider = productservice.NewProductService()
	err = ProductServiceProvider.InitProductComponents(container)
	if err != nil {
		panic(err)
	}

	// LLM service init componets injection
	LLMServiceProvider = llmservice.NewLLMService()
	err = LLMServiceProvider.InitLLMComponents(container)
	if err != nil {
		panic(err)
	}

	return err
}
