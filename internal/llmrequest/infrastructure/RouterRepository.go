package infrastructure

import (
	"PetAi/internal/llmrequest"
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

	OpenAIRepo    llmrequest.LLMRepository `name:"openai"`    // must match OpenAIDiRepositoryName
	AnthropicRepo llmrequest.LLMRepository `name:"anthropic"` // must match AnthropicDiRepositoryName
}

type Router struct {
	repos map[llmrequest.AIApiProvider]llmrequest.LLMRepository
}

func NewLLMRepositoryRouter(p LLMRouterParams) llmrequest.LLMRepositoryRouter {
	return &Router{
		repos: map[llmrequest.AIApiProvider]llmrequest.LLMRepository{
			llmrequest.OpenAi:    p.OpenAIRepo,
			llmrequest.Anthropic: p.AnthropicRepo,
		},
	}
}

func (r *Router) GetRepository(provider llmrequest.AIApiProvider) (llmrequest.LLMRepository, error) {
	repo, ok := r.repos[provider]
	if !ok {
		return nil, errors.New("repository not found for provider: " + string(provider))
	}
	return repo, nil
}
