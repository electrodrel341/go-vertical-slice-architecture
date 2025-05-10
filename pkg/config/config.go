package config

import (
	"PetAi/pkg/apperror"
	"github.com/rs/zerolog/log"
	"os"
	"sync"
)

type Config struct {
	APPConfig   AppConfig
	DBConfig    DBConfig
	ProxyConfig ProxyConfig
	LLMConfig   LLMConfig
	JwtPath     JWTPATH
	// добавьте другие необходимые поля
}

type DBConfig struct {
	User           string
	Password       string
	Name           string
	Host           string
	Port           string
	SSLMode        string
	MigrationsPath string
}

type ProxyConfig struct {
	Url string
}

type AppConfig struct {
	Name     string
	Port     string
	LogLevel string
}

type LLMConfig struct {
	OpenAIAPIKey string
}

type JWTPATH struct {
	PrivateKeyPass string
	PublicKeyPass  string
}

var (
	cfg     *Config
	once    sync.Once
	errInit error
)

func LoadConfig() error {
	once.Do(func() {
		openAIApiKey := os.Getenv("OPENAI_API_KEY")

		if openAIApiKey == "" {
			errInit = apperror.ConfigNotFound(apperror.ErrorNotFoundOpenAiApiKey)
		}

		cfg = &Config{
			LLMConfig: LLMConfig{OpenAIAPIKey: openAIApiKey},
			DBConfig: DBConfig{
				User:           os.Getenv("DB_USER"),
				Password:       os.Getenv("DB_PASSWORD"),
				Name:           os.Getenv("DB_NAME"),
				Host:           os.Getenv("DB_HOST"),
				Port:           os.Getenv("DB_PORT"),
				SSLMode:        os.Getenv("DB_SSL_MODE"),
				MigrationsPath: os.Getenv("MIGRATIONS_PATH"),
			},
			ProxyConfig: ProxyConfig{
				Url: os.Getenv("PROXY_URL"),
			},
			APPConfig: AppConfig{
				Name:     os.Getenv("PROJECT_NAME"),
				Port:     os.Getenv("API_PORT"),
				LogLevel: os.Getenv("LOG_LEVEL"),
			},
			JwtPath: JWTPATH{
				PrivateKeyPass: os.Getenv("JWT_PRIVATE_KEY_PATH"),
				PublicKeyPass:  os.Getenv("JWT_PUBLIC_KEY_PATH"),
			},
		}
	})
	return errInit
}

func Get() *Config {
	if cfg == nil {
		log.Fatal().Msg("Конфигурация не инициализирована. Вызовите Load() в main")
	}
	return cfg
}
