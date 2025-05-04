package logger

import (
	"PetAi/pkg/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func InitLogger() zerolog.Logger {
	level := zerolog.InfoLevel // default

	rawLevel := config.Get().APPConfig.LogLevel // string, например "debug"
	parsedLevel, err := zerolog.ParseLevel(rawLevel)
	if err == nil {
		zerolog.SetGlobalLevel(parsedLevel)
	}
	zerolog.SetGlobalLevel(level)

	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("service", config.Get().APPConfig.Name).
		Logger()

	log.Logger = logger

	return logger
}
