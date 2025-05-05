package infrastructure

import (
	"PetAi/internal/llm"
	"errors"
	"go.uber.org/dig"
)

type DiRepositoryNameType string

const (
	OpenAIDiRepositoryName    DiRepositoryNameType = "openai"
	AnthropicDiRepositoryName DiRepositoryNameType = "anthropic"
)

type LLMRouterParams struct {
	dig.In

	OpenAIRepo    llm.LLMRepository `name:"openai"`    // must match OpenAIDiRepositoryName
	AnthropicRepo llm.LLMRepository `name:"anthropic"` // must match AnthropicDiRepositoryName
}

type Router struct {
	repos map[llm.AIApiProvider]llm.LLMRepository
}

func NewLLMRepositoryRouter(p LLMRouterParams) llm.LLMRepositoryRouter {
	return &Router{
		repos: map[llm.AIApiProvider]llm.LLMRepository{
			llm.OpenAi:    p.OpenAIRepo,
			llm.Anthropic: p.AnthropicRepo,
		},
	}
}

func (r *Router) GetRepository(provider llm.AIApiProvider) (llm.LLMRepository, error) {
	repo, ok := r.repos[provider]
	if !ok {
		return nil, errors.New("repository not found for provider: " + string(provider))
	}
	return repo, nil
}
