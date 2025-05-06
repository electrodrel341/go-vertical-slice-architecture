package main

import (
	"PetAi/pkg/config"
	"PetAi/pkg/injection"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msg("Ошибка загрузки .env файла")
	}

	err = config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Ошибка загрузки конфигурации")
	}

	// prepare all components for dependency injection
	injection.ProvideComponents()

	// initiate service components with dependency injection
	if err := injection.InitComponents(); err != nil {
		log.Fatal().Err(err).Msg("Ошибка инициализации компонентов")
	}

	// run app through dig invoke
	err = injection.InvokeApp()
	if err != nil {
		log.Fatal().Err(err).Msg("Ошибка запуска сервиса")
	}
}
