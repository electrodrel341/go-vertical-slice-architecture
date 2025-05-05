package service

import (
	"PetAi/internal/llm/infrastructure"
	"go.uber.org/dig"
)

// product service instance
type LLMService struct {
	SendLLMRequestServiceProvider *SendLLMRequestService
}

func NewLLMService() *LLMService {
	return &LLMService{}
}

// provide components for injection
func ProvideLLMComponents(c *dig.Container) {
	// repositorory provider injection
	err := c.Provide(infrastructure.NewOpenAIRepository, dig.Name(string(infrastructure.OpenAIDiRepositoryName)))
	if err != nil {
		panic(err)
	}
	err = c.Provide(infrastructure.NewAnthropicRepository, dig.Name(string(infrastructure.AnthropicDiRepositoryName)))
	if err != nil {
		panic(err)
	}

	// Provide router
	err = c.Provide(infrastructure.NewLLMRepositoryRouter)
	if err != nil {
		panic(err)
	}

	// service provider injection
	err = c.Provide(NewSendLLMRequestService)
	if err != nil {
		panic(err)
	}
}

// init service container
func (ps *LLMService) InitLLMComponents(c *dig.Container) error {
	// create product service
	err := c.Invoke(
		func(s *SendLLMRequestService) {
			ps.SendLLMRequestServiceProvider = s
		},
	)

	return err
}
