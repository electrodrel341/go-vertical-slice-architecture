package main

import (
	"PetAi/pkg/config"
	"PetAi/pkg/injection"
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

	// run app through dig invoke
	err = injection.InvokeApp()
	if err != nil {
		log.Fatal(err)
	}
}
