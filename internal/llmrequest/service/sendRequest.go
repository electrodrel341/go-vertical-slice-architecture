package service

import (
	"PetAi/internal/llmrequest"
	"PetAi/pkg/apperror"
	"github.com/rs/zerolog"
)

type SendLLMRequestService struct {
	router llmrequest.LLMRepositoryRouter
	logger zerolog.Logger
}

func NewSendLLMRequestService(
	router llmrequest.LLMRepositoryRouter,
	logger zerolog.Logger,
) *SendLLMRequestService {
	return &SendLLMRequestService{
		router: router,
		logger: logger,
	}
}

func (service *SendLLMRequestService) SendRequest(p *llmrequest.LLMRequest) (string, error) {

	repo, err := service.router.GetRepository(p.AIApiProvider)
	response, err := repo.SendRequest(&p.Promt)

	if err != nil {
		service.logger.Error().Str("provider", string(p.AIApiProvider)).Msg("Error sending LLM request")
		err := apperror.InternalServerError(err)
		return "", err
	}

	return response, nil
}
